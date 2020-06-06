import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:fixnum/fixnum.dart';
import '../services/chat_service.dart';
import 'message_page.dart';
import 'create_chatroom.dart';

class LoginPage extends StatefulWidget {
  final FirebaseUser user;
  LoginPage(this.user);
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  String name;
  String videoURL;
  Int64 chatroom;
  TextEditingController controller;
  TextEditingController chatroomController;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    name = widget.user.displayName;
    chatroom = Int64(new DateTime.now().millisecondsSinceEpoch);
    controller = TextEditingController(text: name);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Join Chatroom"),
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
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.all(Radius.circular(20.0))),
                    labelText: "Enter Your Chatroom ID",
                  ),
                  onChanged: (value) {
                    chatroom = Int64(int.tryParse(value));
                    print(chatroom);
                  }),
              SizedBox(
                height: 10.0,
              ),
              TextField(
                  keyboardType: TextInputType.number,
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
                          MaterialPageRoute(builder: (context) {
                            return MessagePage(ChatService(name, chatroom),
                                chatroom, false, videoURL);
                          }),
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
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          Navigator.push(
            context,
            MaterialPageRoute(builder: (context) {
              return CreateChatroomPage(widget.user);
            }),
          );
        },
        label: Text("Create a new room"),
        backgroundColor: Colors.red,
        icon: Icon(Icons.add_circle_outline),
      ),
    );
  }
}
