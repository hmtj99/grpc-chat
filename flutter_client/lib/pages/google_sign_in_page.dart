import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'login_page.dart';
import 'package:modal_progress_hud/modal_progress_hud.dart';
import 'package:flutter_auth_buttons/flutter_auth_buttons.dart';

class GoogleSignInPage extends StatefulWidget {
  @override
  _GoogleSignInPageState createState() => _GoogleSignInPageState();
}

class _GoogleSignInPageState extends State<GoogleSignInPage> {
  final GoogleSignIn _googleSignIn = GoogleSignIn();
  final FirebaseAuth _auth = FirebaseAuth.instance;
  FirebaseUser loggedInUser;
  bool shouldSpin = false;

  Future<FirebaseUser> _handleSignIn() async {
    final GoogleSignInAccount googleUser = await _googleSignIn.signIn();
    final GoogleSignInAuthentication googleAuth =
        await googleUser.authentication;

    final AuthCredential credential = GoogleAuthProvider.getCredential(
      accessToken: googleAuth.accessToken,
      idToken: googleAuth.idToken,
    );

    final FirebaseUser user =
        (await _auth.signInWithCredential(credential)).user;
    print("signed in " + user.email);
    return user;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Youtube Rooms"),
      ),
      body: ModalProgressHUD(
        inAsyncCall: shouldSpin,
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Image.asset(
                "images/logo_red2.png",
              ),
              GoogleSignInButton(
                onPressed: () {
                  _handleSignIn().then((FirebaseUser user) {
                    setState(() {
                      shouldSpin = true;
                    });
                    loggedInUser = user;
                    Navigator.push(context,
                        MaterialPageRoute(builder: (context) {
                      return LoginPage(loggedInUser);
                    }));
                  }).catchError((e) => print(e));
                },
                darkMode: true,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
