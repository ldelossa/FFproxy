package server

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	pb "github.com/ldelossa/ffproxy/proto"
)

// FFproxy implements our GRPC server and methods for managing stream abstractions.
// implements core methods for bridging http -> grpc and back
type FFproxy struct {
	// lock protecting stream map
	sl sync.RWMutex
	// stream map associates hostnames with streams
	streamMap map[string]*Stream
}

// DoProxy is a bridge between an http handler and the grpc client stream.
// takes *http.Request, finds associated stream if any, translates http request to Proto
// and makes a blocking call to our stream abstraction's send function. see stream.go
// for further explanation of streama abstraction
func (f *FFproxy) DoProxy(req *http.Request) (*http.Response, error) {
	// lookup stream associated with host
	f.sl.RLock()
	stream, ok := f.streamMap[req.Host]
	f.sl.RUnlock()

	if !ok {
		return nil, fmt.Errorf("failed to find stream for hostname: %s", req.Host)
	}

	// write http request to byte array
	buf := &bytes.Buffer{}
	err := req.Write(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to write http response into buffer: %s", err)
	}

	// create proto
	msg := &pb.ServerClient{
		Servermsg: &pb.ServerClient_Httpreq{
			Httpreq: &pb.HTTPRequest{
				Request: buf.Bytes(),
			},
		},
	}

	// block on stream.Send() and wait for response
	respmsg, err := stream.Send(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to send proto on stream: %s", err)
	}

	// unpack returned proto to *http.Response
	switch msg := respmsg.Clientmsg.(type) {
	case *pb.ClientServer_Httpresp:
		buf := bytes.NewBuffer(msg.Httpresp.Response)
		r := bufio.NewReader(buf)
		resphttp, err := http.ReadResponse(r, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to read proto into http response: %s", err)
		}

		return resphttp, nil
	default:
		return nil, fmt.Errorf("received unhandled message type as response %T", msg)
	}

}

// Stream method implements our grpc bidirectional server. this method is called on every
// client stream connection. we handle InitRequest messages which act as a handshake with our client.
// stream abstractions are assigned to the server's stream map on successful handshake.
func (f *FFproxy) Stream(stream pb.FFProxy_StreamServer) error {

	// receive first proto from stream
	req, err := stream.Recv()
	if err != nil {
		log.Printf("failed to receive handshake for new client: %s", err)
		return err
	}

	// type switch on received message. we expect this to be a InitRequest message. if it is not return error
	switch m := req.Clientmsg.(type) {
	case *pb.ClientServer_Initreq:
		// create UUID string to be used for hostname
		hostname := uuid.New().String()

		// create ServerClient_InitResp to tell client it's hostname
		initResp := &pb.ServerClient{
			Servermsg: &pb.ServerClient_Initresp{
				Initresp: &pb.InitResponse{
					Hostname: hostname,
				},
			},
		}

		// create our stream abstraction. all subsequent sending and receiving on the grpc stream will be done by our self managing abstraction
		st := NewStream(stream, cancelFunc(f))

		// register stream abstraction with server
		f.sl.Lock()
		f.streamMap[hostname] = st
		f.sl.Unlock()

		// return initResp to client, any subsequent receives are handled by self managing stream abstraction
		err := stream.Send(initResp)
		if err != nil {
			log.Printf("failed to send init response to client: %s", err)
			// TODO: should we remove stream from map?
			return fmt.Errorf("failed to send init response to client: %s", err)
		}

		return nil
	default:
		e := fmt.Sprintf("expected message type ClientServer_Initreq but got %T", m)
		log.Printf(e)
		return fmt.Errorf(e)
	}
}

// removeStream removes a stream from our server's stream map
func (f *FFproxy) removeStream(hostname string) error {
	// get readlock to see if hostname exists in map
	f.sl.RLock()
	_, ok := f.streamMap[hostname]
	f.sl.RUnlock()

	if ok {
		// take lock
		f.sl.Lock()
		delete(f.streamMap, hostname)
		f.sl.Unlock()
		return nil
	}

	return fmt.Errorf("failed to find stream %s", hostname)
}

// cancelFunc is passed to our stream abstraction and is called when the stream receives
// an error on the underlying grpc stream.
func cancelFunc(f *FFproxy) func(streamID string) {
	return func(streamID string) {
		// remove stream from server's stream map
		err := f.removeStream(streamID)
		if err != nil {
			log.Printf("failed to remove stream %s from server's map: %s", streamID, err)
		}
		return
	}
}
