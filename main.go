package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gritt/go-protocol-buffer/src/pkg/domain/types"
	"google.golang.org/protobuf/proto"
)

func main() {
	testWriteContact()
	testWriteAddressBook()
	testReadAddressBook()
}

func testWriteContact() {
	person := getFakePerson()

	data, err := proto.Marshal(&person)
	if err != nil {
		fmt.Println("marshal err: ", err)
	}

	if err := ioutil.WriteFile("./contact.txt", data, fs.ModePerm); err != nil {
		fmt.Println("write file err: ", err)
	}
}

func testWriteAddressBook() {
	addressBook := types.AddressBook{}
	for i := 0; i < 100; i++ {
		person := getFakePerson()
		addressBook.Contact = append(addressBook.Contact, &person)
	}

	data, err := proto.Marshal(&addressBook)
	if err != nil {
		fmt.Println("marshal err: ", err)
	}

	if err := ioutil.WriteFile("./addressbook.txt", data, fs.ModePerm); err != nil {
		fmt.Println("write file err: ", err)
	}
}

func testReadAddressBook() {
	data, err := ioutil.ReadFile("./addressbook.txt")
	if err != nil {
		fmt.Println("read file err:", err)
	}

	addressBook := &types.AddressBook{}
	if err := proto.Unmarshal(data, addressBook); err != nil {
		fmt.Println("unmarshal err:", err)
	}

	fmt.Println("== ADDRESS BOOK ==")
	fmt.Println(addressBook)
}

func getFakePerson() types.Person {
	return types.Person{
		Id:    getFakeID(),
		Name:  getFakeName(),
		Phone: getFakePhone(),
	}
}

func getFakeID() int32 {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	return int32(id)
}

func getFakePhone() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

func getFakeName() string {
	return uuid.NewString()
}
