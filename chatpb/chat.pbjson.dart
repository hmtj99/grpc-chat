///
//  Generated code. Do not modify.
//  source: chat.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const User$json = const {
  '1': 'User',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'chatRoomNo', '3': 3, '4': 1, '5': 3, '10': 'chatRoomNo'},
  ],
};

const Message$json = const {
  '1': 'Message',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'content', '3': 2, '4': 1, '5': 9, '10': 'content'},
    const {'1': 'timestamp', '3': 3, '4': 1, '5': 9, '10': 'timestamp'},
    const {'1': 'chatRoom', '3': 4, '4': 1, '5': 3, '10': 'chatRoom'},
    const {'1': 'isYTLink', '3': 5, '4': 1, '5': 8, '10': 'isYTLink'},
  ],
};

const Connect$json = const {
  '1': 'Connect',
  '2': const [
    const {'1': 'user', '3': 1, '4': 1, '5': 11, '6': '.chat.User', '10': 'user'},
    const {'1': 'active', '3': 2, '4': 1, '5': 8, '10': 'active'},
    const {'1': 'wantsNewRoom', '3': 3, '4': 1, '5': 8, '10': 'wantsNewRoom'},
  ],
};

const Close$json = const {
  '1': 'Close',
};

const BroadcastServiceBase$json = const {
  '1': 'Broadcast',
  '2': const [
    const {'1': 'CreateStream', '2': '.chat.Connect', '3': '.chat.Message', '6': true},
    const {'1': 'BroadcastMessage', '2': '.chat.Message', '3': '.chat.Close'},
  ],
};

const BroadcastServiceBase$messageJson = const {
  '.chat.Connect': Connect$json,
  '.chat.User': User$json,
  '.chat.Message': Message$json,
  '.chat.Close': Close$json,
};

