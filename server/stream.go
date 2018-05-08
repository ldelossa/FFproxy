package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	pb "github.com/ldelossa/ffproxy/proto"
)

// Stream abstracts the underlying grpc stream object.
// Stream objects are self managing. on client disconnect they will return from their Rx and Tx loops
// TODO: gracefully exit from Rx and Tx loops on client disconnection
// and call cancelFunc if not nil.
type Stream struct {
	id string
	// the underlying grpc stream
	stream pb.FFProxy_StreamServer
	// channel to serialize access to underlying grpc stream end method (not go routine safe)
	sendChan chan *pb.ServerClient
	// protect access to requestID map
	rl sync.RWMutex
	// map which associates request ID to
	requestIDMap map[string]chan<- *pb.ClientServer
	// cancelFunc is called when an error is encountered on the underyling grpc stream
	// any cleanup logic can be placed here. Nil is an acceptable value
	cancelFunc func(streamID string)
}

// NewStream is the constructor method for our stream abstraction
// receives a grpc stream object and a cancelFunc called on the receive of an error on the grpc stream
func NewStream(stream pb.FFProxy_StreamServer, cancelFunc func(streamID string)) *Stream {
	// create internal send channel
	c := make(chan *pb.ServerClient, 1024)

	// create requestID->responseChan map
	m := make(map[string]chan<- *pb.ClientServer, 0)

	// create stream
	s := &Stream{
		stream:       stream,
		sendChan:     c,
		requestIDMap: m,
		cancelFunc:   cancelFunc,
	}

	// start Rx and Tx loops as go routines
	go s.gRPCRxLoop()
	go s.gRPCTxLoop()

	return s
}

// Send is the external API for sending a ServerClient message and receiving
// it's response. Callers can block on this method waiting for a response from the stream
// TODO: plumb context thru this
func (s *Stream) Send(msg *pb.ServerClient) (*pb.ClientServer, error) {
	// create request ID
	id := uuid.New().String()

	// add requestID to proto
	msg.Servermsg.(*pb.ServerClient_Httpreq).Httpreq.RequestUUID = id

	// create channel to block on until response
	respChan := make(chan *pb.ClientServer)

	// associate respChan with requestID in map
	s.rl.Lock()
	s.requestIDMap[id] = respChan
	s.rl.Unlock()

	// send message to sendChan
	select {
	case s.sendChan <- msg:
	default:
		return nil, fmt.Errorf("failed to send proto to underlying stream")
	}

	// block on response channel until resp is received
	csm := <-respChan

	// return responded ClientServer message to caller
	return csm, nil
}

// onRecv is called on receiving a *pb.ClientServer_Httpresp and the message to the curently blocking Send() method call.
func (s *Stream) onRecv(msg *pb.ClientServer) {
	// extract requestUUID from msg we. know this will pass. see gRPCRxLoop and usage of this function
	reqID := msg.Clientmsg.(*pb.ClientServer_Httpresp).Httpresp.RequestUUID
	if reqID == "" {
		log.Printf("recieved proto with blank request ID. no further processing")
		return
	}

	// lookup ID
	s.rl.RLock()
	c, ok := s.requestIDMap[reqID]
	s.rl.Unlock()

	if !ok {
		log.Printf("received unknowing request ID %s. no channel to respond to")
		return
	}

	// send response to channel. this will unblock the caller waiting on s.Send() for this request ID
	select {
	case c <- msg:
		log.Printf("returned proto for request ID %s")
	default:
		log.Printf("failed to send received proto upstream to response channel")
	}
}

// gRPCRxLoop begins a receive loop on the underlying grpc stream. on receipt of a ClientServer_Httpresp proto we call
// s.onRecv for further processing
func (s *Stream) gRPCRxLoop() {
	for {
		msg, err := s.stream.Recv()
		// handle stream error
		if err != nil {
			log.Printf("received error on grpc stream. closing stream: %s", err)
			// call cancel func to perform any operations involving server cleanup
			// of this stream
			if s.cancelFunc != nil {
				s.cancelFunc(s.id)
			}
			return
		}

		// type switch on message
		switch m := msg.Clientmsg.(type) {
		case *pb.ClientServer_Httpresp:
			// call onRecv as a go routine so processing is faster then blocking here
			go s.onRecv(msg)
		default:
			log.Printf("received unhandle message type %T", m)
			continue
		}
	}
}

// gRPCTxLoop begins listening on the internal sendChan for external api calls trying to enqueue a message.
// calls send on the underying grpc stream. serializes access to grpc Send method which is not go routine safe
// TODO: work in elegant cancelation of this channel. Context ? Nil channels ?
func (s *Stream) gRPCTxLoop() {
	for {
		msg := <-s.sendChan
		err := s.stream.Send(msg)
		if err != nil {
			log.Printf("failed to send message to grpc stream: %s", err)
			continue
		}
	}
}
