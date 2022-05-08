package sendTextMessages

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"sub/modelWA"
	"sub/system/binary/proto"

	"sub/system"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
)

func SendText(req modelWA.User) modelWA.ResGlobal {
	var result modelWA.ResGlobal
	var isiResult modelWA.ResSendWA
	t := time.Now()
	fmt.Println("lewat sendText.SendText____________:")
	//create new WhatsApp connection
	wac, err := system.NewConn(5 * time.Second)
	if err != nil {
		result.Status = "31"
		result.StatusDateTime = t
		result.StatusDesc = err.Error()
		result.Result = isiResult
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return result
	}

	err = login(wac)
	if err != nil {
		result.Status = "31"
		result.StatusDateTime = t
		result.StatusDesc = err.Error()
		result.Result = isiResult
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return result
	}

	<-time.After(3 * time.Second)

	previousMessage := "😘"
	quotedMessage := proto.Message{
		Conversation: &previousMessage,
	}

	ContextInfo := system.ContextInfo{
		QuotedMessage:   &quotedMessage,
		QuotedMessageID: "",
		Participant:     "", //Who sent the original message
	}
	nomer := "62" + req.No
	msg := system.TextMessage{
		Info: system.MessageInfo{
			RemoteJid: nomer + "@s.whatsapp.net",
		},
		ContextInfo: ContextInfo,
		Text:        req.Text,
	}

	msgId, err := wac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
		// os.Exit(1)
		result.Status = "81"
		result.StatusDateTime = t
		result.StatusDesc = err.Error()
		return result
	} else {
		fmt.Println("Message Sent -> ID : " + msgId)
	}

	// isiResult.ID = msgId
	isiResult.To = nomer
	isiResult.Text = req.Text
	result.Status = "00"
	result.StatusDateTime = t
	result.StatusDesc = "success"
	result.Result = isiResult
	return result
}

func login(wac *system.Conn) error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v\n", err)
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v\n", err)
		}
	}

	//save session
	err = writeSession(session)
	if err != nil {
		return fmt.Errorf("error saving session: %v\n", err)
	}
	return nil
}

func readSession() (system.Session, error) {
	session := system.Session{}
	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session system.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}
