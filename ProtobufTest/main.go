package main

import (
	"ProtobufTest/example"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	test := &example.TestMessage{
		ClientName:  proto.String("client"),
		ClientId:    proto.Int32(1),
		Description: proto.String("hello"),
	}

	ItemTypeX := example.TestMessage_TypeX

	MsgItem1 := &example.TestMessage_MsgItem{
		Id:        proto.Int32(1),
		ItemName:  proto.String("FirstItemName"),
		ItemValue: proto.Int32(222),
		ItemType:  ItemTypeX.Enum(),
	}

	test.MessageItems = append(test.MessageItems, MsgItem1)

	ItemTypeY := example.TestMessage_TypeY

	MsgItem2 := &example.TestMessage_MsgItem{
		Id:        proto.Int32(2),
		ItemName:  proto.String("SecondItemName"),
		ItemValue: proto.Int32(333),
		ItemType:  ItemTypeY.Enum(),
	}

	test.MessageItems = append(test.MessageItems, MsgItem2)

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error\n")
		return
	}

	newTest := &example.TestMessage{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("unmarshaling error\n")
		return
	}

	fmt.Printf("%s\n", newTest.GetClientName())
	fmt.Printf("%v\n", newTest.GetClientId())
	fmt.Printf("%s\n", newTest.GetDescription())

	NewMessageItem := newTest.GetMessageItems()
	for i := 0; i < len(NewMessageItem); i++ {
		fmt.Printf("%v	%s, %v, %v\n", NewMessageItem[i].GetId(), NewMessageItem[i].GetItemName(),
			NewMessageItem[i].GetItemValue(), NewMessageItem[i].GetItemType())
	}
}
