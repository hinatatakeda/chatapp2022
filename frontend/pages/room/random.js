// import {useState, useEffect} from 'react';
// import axios from 'axios'
// import Link from 'next/link'
// import styles from '../../styles/Home.module.css'


// const Room = () =>{
//   const [chats, setChats] = useState([]);
//   const [text, setText] = useState("");
//   const [userID, setUserID] = useState("1");
//   // chat dataを取得
//   useEffect(() => {
//     axios.get('http://localhost:8080/chat?roomID=2')
//       .then((res) => {
//         console.log(res.data);
//         setChats(res.data);
//       })
//       .catch((err) => {
//         console.log(err);
//       })
//   }, []);


//   const postChat = (text) => {
//     if (text===""){
//       return;
//     }
//     axios.post(`http://localhost:8080/chat/post?roomID=1&userID=${userID}&chatText=${text}`)
//       .then((res) => {
//         console.log(res.data);
//         setText("");
//       })
//       .catch((err) => {
//         console.log(err);
//         setText("");
//       })
//     }

//   // chat画面を表示
//   return (
//     <div>
//       <div>
//       <div className={styles.generalhead}>
//       <h1>Random</h1>
//       <div>
//       <Link href="./..">
//         <a>HOME</a>
//       </Link>
//       </div>
//       </div> 
//       </div> 


//       <div className={styles.contanr}>



// <button onClick={(userID)=>{
//   if (userID==="1"){
//     setUserID("2");
//   } else {
//     setUserID("1");
//   }
// }}>ユーザーを切り替える</button>
// </div>

        
        





//     {chats.map((chat) => {
//     return (
//       <div className={styles.sendgeneral} >
//         {chat.isEditted ? <p>編集済み</p> : null}
//         {chat.UserID==="1"? <p>送信者：名探偵コナン</p>:<p>送信者：おひな。</p> }
//         <p>{chat.Text}</p>
//       </div>
//     )
//   })}



//   <div className={styles.send}>
//   <input type="text" value={text} onChange={(e) => setText(e.target.value)} />
//   <button onClick={()=>{
//     postChat(text, 2);
//   }}
//   >送信</button> </div> 

// </div> 
//   )
// }




// export default Room;


import {useState, useEffect} from 'react';
import axios from 'axios'
import Link from 'next/link'

import styles from '../../styles/Home.module.css'

const Random = () => {
  const [chats, setChats] = useState([]);
  const [text, setText] = useState("");
  const [userID, setUserID] = useState("1");
  const [isEditting, setIsEditting] = useState(false);
  const [edittingChatID, setEdittingChatID] = useState("");
  // chat dataを取得
  useEffect(() => {
    axios.get('http://localhost:8080/chat?roomID=2')
      .then((res) => {
        console.log(res.data);
        setChats(res.data);
      })
      .catch((err) => {
        console.log(err);
      })
  }, []);

  const postChat = (text) => {
    if (text===""){
      return;
    }
    axios.post(`http://localhost:8080/chat?roomID=2&userID=${userID}&chatText=${text}`)
      .then((res) => {
        console.log(res.data);
        setText("");
      })
      .catch((err) => {
        console.log(err);
        setText("");
      })
    }


  // {/* delete */}
  const deleteChat = (chatID) => {
    axios.delete(`http://localhost:8080/chat?chatID=${chatID}`)
      .then((res) => {
        console.log(res.data);
      })
      .catch((err) => {
        console.log(err);
      })
    }


   // {/* edit*/}
   const editChat = (chatID, text) => {
    if (text===""){
      return;
    }
    axios.put(`http://localhost:8080/chat?chatID=${chatID}&chatText=${text}`)
      .then((res) => {
        console.log(res.data);
        setIsEditting(false);
        setEdittingChatID("");
        setText("");
      })
      .catch((err) => {
        console.log(err);
        setIsEditting(false);
        setEdittingChatID("");
        setText("");
      })
    }










  // chat画面を表示
  return (
    <div>
      <div>
      <div className={styles.leftmenu}>


      <div>  </div>

      <div>
      <Link href="./..">
        <a>RETURN HOME</a>
      </Link>
      </div>



      <div>
      {/* <div className={styles.leftmenu}> USER </div> */}

      <div>ユーザー選択</div>
      <button onClick={()=>{
      setUserID("1")
        }}>名探偵コナン</button>

      <button onClick={()=>{
      setUserID("2")
        }}>おひな。</button>

      </div>
      



  






      <div className={styles.contanr}>



        <button onClick={(userID)=>{
          if (userID==="1"){
            setUserID("2");
          } else {
            setUserID("1");
          }
        }}>ユーザーを切り替える</button>
      </div>

        {chats.map((chat) => {
          return (
            <div className={styles.sendgeneral} >
               <button onClick={()=>{
                setIsEditting(true);
                setEdittingChatID(chat.ChatID);
                setText(chat.Text);
               }}>
                編集する
              </button>
              <button onClick={()=>{
                deleteChat(chat.ChatID);
              }}
               >削除
               </button>
              {chat.IsEditted ? <p>編集済み</p> : null}
              {chat.UserID==="1"? <p>送信者：hinata</p>:<p>送信者：takeda</p> }
              <p>{chat.Text}</p>
            </div>
          )
        })}
        
        </div> 
        <div className={styles.send}>
        <input type="text" value={text} onChange={(e) => setText(e.target.value)} />
        {isEditting ? (<button onClick={()=>{
          editChat(edittingChatID, text);
        }}
        >編集</button>):(<button onClick={()=>{
          postChat(text, 2);
        }}
        >送信</button>)}
        </div> 





      </div>  
    
  )


}
export default Random ;