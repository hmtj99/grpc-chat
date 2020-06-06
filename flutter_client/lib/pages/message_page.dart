import 'package:flutter/material.dart';
import 'package:fixnum/fixnum.dart';
import 'package:youtube_player_flutter/youtube_player_flutter.dart';
import '../services/chat_service.dart';
import '../protos/chat.pb.dart';
import 'dart:async';

class MessagePage extends StatefulWidget {
  final ChatService service;
  final Int64 chatroomid;
  final bool wantsNewRoom;
  final String videoURL;
  MessagePage(this.service, this.chatroomid, this.wantsNewRoom, this.videoURL);
  @override
  _MessagePageState createState() => _MessagePageState();
}

class _MessagePageState extends State<MessagePage> {
  @override
  void initState() {
    super.initState();
    controller = TextEditingController(
        text: "https://www.youtube.com/watch?v=59AYXzCa-Cs");
    messages = Set();
    scrollController = ScrollController();
    _ytcontroller = YoutubePlayerController(
        initialVideoId: YoutubePlayer.convertUrlToId(widget.videoURL),
        flags: YoutubePlayerFlags(autoPlay: false));
  }

  Stream<List<int>> stream;
  StreamSubscription<List<int>> listener;
  Set<Message> messages;
  TextEditingController controller;
  ScrollController scrollController;
  YoutubePlayerController _ytcontroller;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Chatroom - ' + widget.chatroomid.toString()),
        actions: <Widget>[
          IconButton(
              icon: Icon(
                Icons.add,
                color: Colors.white,
              ),
              onPressed: () {
                widget.service.sendMessage(controller.text, true);
                controller.clear();
              })
        ],
      ),
      body: Center(
        child: Column(
          children: <Widget>[
            YoutubePlayer(
              controller: _ytcontroller,
              showVideoProgressIndicator: true,
              progressIndicatorColor: Colors.amber,
              progressColors: ProgressBarColors(
                  playedColor: Colors.amber, handleColor: Colors.amberAccent),
            ),
            Flexible(
              child: StreamBuilder<Message>(
                stream: widget.service
                    .recieveMessage(wantsNewRoom: widget.wantsNewRoom),
                builder: (context, snapshot) {
                  if (snapshot.hasError) {
                    print("There was an error!!!");
                    return Center(
                      child: AlertDialog(
                        title: Text("Error"),
                        content: Text(snapshot.error),
                        actions: <Widget>[
                          FlatButton(
                              onPressed: () {
                                listener.cancel();
                                Navigator.pop(context);
                              },
                              child: Text("OK"))
                        ],
                      ),
                    );
                  }
                  if (!snapshot.hasData) {
                    return Center(child: Text("Start by typing a message"));
                  }
                  if (snapshot.data.isYTLink == false) {
                    messages.add(snapshot.data);
                  } else {
                    String videoID =
                        YoutubePlayer.convertUrlToId(snapshot.data.content);
                    _ytcontroller.load(videoID);
                  }
                  return ListView(
                    controller: scrollController,
                    children: messages
                        .map(
                          (msg) => MessageBubble(
                            senderName: msg.id,
                            content: msg.content,
                            timestamp: msg.timestamp,
                            isMe: msg.id == widget.service.user.name,
                          ),
                        )
                        .toList(),
                  );
                },
              ),
            ),
            Row(
              children: <Widget>[
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.fromLTRB(8.0, 8.0, 8.0, 3.0),
                    child: TextField(
                      controller: controller,
                      decoration: InputDecoration(
                          border: OutlineInputBorder(
                              borderRadius:
                                  BorderRadius.all(Radius.circular(20.0))),
                          labelText: "Type a message"),
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(5.0),
                  child: FloatingActionButton(
                    onPressed: () {
                      widget.service.sendMessage(controller.text, false);
                      controller.clear();
                      scrollController
                          .jumpTo(scrollController.position.maxScrollExtent);
                    },
                    child: Icon(Icons.send),
                    backgroundColor: Colors.redAccent,
                  ),
                )
              ],
            )
          ],
        ),
      ),
    );
  }
}

class MessageBubble extends StatelessWidget {
  final String senderName;
  final String content;
  final String timestamp;
  final bool isMe;

  MessageBubble({this.senderName, this.content, this.timestamp, this.isMe});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: Column(
        crossAxisAlignment:
            isMe ? CrossAxisAlignment.end : CrossAxisAlignment.start,
        children: <Widget>[
          Text(
            senderName,
            style: TextStyle(
              color: Colors.black54,
              fontSize: 12.0,
            ),
          ),
          Material(
            elevation: 5.0,
            borderRadius: isMe
                ? BorderRadius.only(
                    topLeft: Radius.circular(30.0),
                    bottomLeft: Radius.circular(30.0),
                    bottomRight: Radius.circular(30.0))
                : BorderRadius.only(
                    topRight: Radius.circular(30.0),
                    bottomLeft: Radius.circular(30.0),
                    bottomRight: Radius.circular(30.0)),
            color: isMe ? Colors.redAccent : Colors.white,
            child: Padding(
              padding:
                  const EdgeInsets.symmetric(vertical: 10.0, horizontal: 20.0),
              child: Text(
                '$content',
                style: TextStyle(
                  fontSize: 15.0,
                  color: isMe ? Colors.white : Colors.black54,
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}
