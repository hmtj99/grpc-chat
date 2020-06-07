import 'dart:async';

import 'package:fixnum/fixnum.dart';
import 'package:flutter_client/protos/chat.pb.dart';
import 'package:flutter_client/protos/chat.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:crypto/crypto.dart';
import 'dart:convert';

class ChatService {
  User user = User();
  static BroadcastClient client;

  ChatService(String name, Int64 chatroom) {
    user
      ..clearName()
      ..name = name
      ..clearId()
      ..id = sha256.convert(utf8.encode(user.name)).toString()
      ..clearChatRoomNo()
      ..chatRoomNo = chatroom;

    client = BroadcastClient(
      ClientChannel(
        "34.220.130.254",
        port: 80,
        options: ChannelOptions(credentials: ChannelCredentials.insecure()),
      ),
    );
  }

  Future<Close> sendMessage(String body, bool isYTLink) async {
    print("Send Message Called");
    return client.broadcastMessage(Message()
      ..id = user.name
      ..content = body
      ..timestamp = DateTime.now().toIso8601String()
      ..chatRoom = user.chatRoomNo
      ..isYTLink = isYTLink);
  }

  Stream<Message> recieveMessage({bool wantsNewRoom}) async* {
    print("RecieveMsg Called");
    print(user);
    Connect conn = Connect()
      ..user = user
      ..active = true
      ..wantsNewRoom = wantsNewRoom;

    StreamController<Message> messageController =
        StreamController<Message>.broadcast();
    messageController.addStream(client.createStream(conn));

    // try {
    await for (var msg in messageController.stream) {
      print("in the loop");
      yield msg;
    }
    // } catch (e) {
    //   messageController.sink.addError(e);
    //   await for (var msg in messageController.stream) {
    //     yield msg;
    //   }
    // }
  }
}
