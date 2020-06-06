import 'package:flutter/material.dart';
import 'package:flutter_client/pages/login_page.dart';
import 'pages/google_sign_in_page.dart';
import 'package:firebase_auth/firebase_auth.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final FirebaseAuth _auth = FirebaseAuth.instance;
  FirebaseUser user;

  Future getCurrentUser() async {
    user = await _auth.currentUser();
    return user;
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "GRPC Chat",
      theme: ThemeData(
          appBarTheme: AppBarTheme(
        color: Colors.red,
      )),
      home: user == null ? GoogleSignInPage() : LoginPage(user),
    );
  }
}
