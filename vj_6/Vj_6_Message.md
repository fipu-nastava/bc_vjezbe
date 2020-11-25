<table style="caret-color: #000000; font-family: Georgia;" border="0" cellspacing="0" cellpadding="0" >
            <tbody>
              <tr>
                <td valign="center">
                  <a id="logo_a" href="https://fipu.unipu.hr"><img id="logo_img"  src="https://www.unipu.hr/_download/repository/FIPU_horiz_kolor_HR.png" alt="logotip FIPU" title="Fakultet informatike u Puli"></a> 								 </td>
              </tr>
  </tbody>
</table>



<img src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" alt="Golang" style="zoom:25%;" />



# Message

- Napravite strukturu 

  ```go
  Message {
    Sender  string
    Msg    string
  }
  ```

- Napišite metode za strukturu `Message` 

  ```go
  // Računa i vraća Hash Message-a (Hint: pozvati metodu Serialize za pretvorbu u []byte)
  func (m *Message) Hash() Hash
  
  // Pretvara cijeli Message u niz byteova
  func (m *Message) Serialize() []byte
  
  // Pretvara serijaliziran Message u strukturu **(Serijalizacija != Hash)**
  func Deserialize(reader io.Reader) (*Message, error)
  ```

- ```go
  func main() { // Main funkcija za testiranje strukture
     msg := &Message{"Me", "Test"}
  
     buffer := &bytes.Buffer{}
     rw := bufio.NewReadWriter(bufio.NewReader(buffer), bufio.NewWriter(buffer))
  
     rw.Write(msg.Serialize())
     rw.Flush()
  
     msg, err := Deserialize(rw)
     if err != nil {
        panic(err)
     }
    
     fmt.Println(msg)
  }
  ```

