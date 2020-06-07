import 'package:flutter/material.dart';
import 'package:fixnum/fixnum.dart';
import 'message_page.dart';
import '../services/chat_service.dart';
import 'package:firebase_auth/firebase_auth.dart';

class CreateChatroomPage extends StatefulWidget {
  final FirebaseUser user;
  CreateChatroomPage(this.user);
  @override
  _CreateChatroomPageState createState() => _CreateChatroomPageState();
}

class _CreateChatroomPageState extends State<CreateChatroomPage> {
  String name;
  String videoURL;
  Int64 chatroom;
  TextEditingController controller;
  TextEditingController controller2;
  TextEditingController chatroomController;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    chatroom = Int64(new DateTime.now().millisecondsSinceEpoch);
    controller = TextEditingController(text: widget.user.displayName);
    controller2 = TextEditingController();
    chatroomController = TextEditingController(text: chatroom.toString());
    name = widget.user.displayName;
    videoURL = "https://www.youtube.com/watch?v=4uMchksQJng";
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Create a new Chatroom"),
        backgroundColor: Colors.red,
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              TextField(
                  keyboardType: TextInputType.text,
                  controller: controller,
                  decoration: InputDecoration(
                    labelText: "Enter Your Name",
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.all(Radius.circular(20.0))),
                  ),
                  onChanged: (value) {
                    name = value;
                    print(name);
                  }),
              SizedBox(
                height: 10.0,
              ),
              TextField(
                  keyboardType: TextInputType.number,
                  controller: chatroomController,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.all(Radius.circular(20.0))),
                    labelText: "Enter New Chatroom ID",
                  ),
                  onChanged: (value) {
                    chatroom = Int64(int.tryParse(value));
                    print(chatroom);
                  }),
              SizedBox(
                height: 20.0,
              ),
              TextField(
                  autofocus: false,
                  keyboardType: TextInputType.number,
                  controller: controller2,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.all(Radius.circular(20.0))),
                    labelText: "Enter Youtube Video URL",
                  ),
                  onChanged: (value) {
                    videoURL = value;
                    print(videoURL);
                  }),
              SizedBox(
                height: 20.0,
              ),
              Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: <Widget>[
                    MaterialButton(
                      elevation: 5.0,
                      height: 40.0,
                      minWidth: 100.0,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(10.0),
                      ),
                      color: Colors.red,
                      onPressed: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => MessagePage(
                                ChatService(name, chatroom),
                                chatroom,
                                true,
                                videoURL),
                          ),
                        );
                      },
                      child: Text(
                        "Submit",
                        style: TextStyle(color: Colors.white),
                      ),
                    ),
                  ])
            ],
          ),
        ),
      ),
    );
  }
}
