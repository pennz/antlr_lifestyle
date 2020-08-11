import 'dart:async';
import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;

import 'package:flutter/material.dart';
import 'package:english_words/english_words.dart'; // 新增了这一行
import 'package:flutter_svg/flutter_svg.dart';

void main() => runApp(new MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    String appTitle = 'Startup Name Generator';

    return new MaterialApp(
      title: appTitle,
      theme: new ThemeData(
        // 新增代码开始...
        primaryColor: Colors.white,
      ), // ... 代码新增结束
      //home: new RandomWords(),
      home: MyHomePage(title: appTitle),
    );
  }
}

class RandomWords extends StatefulWidget {
  @override
  RandomWordsState createState() => new RandomWordsState();
}

class RandomWordsState extends State<RandomWords> {
  final List<WordPair> _suggestions = <WordPair>[];
  final Set<WordPair> _saved = new Set<WordPair>(); // 新增本行
  final TextStyle _biggerFont = const TextStyle(fontSize: 18.0);

  Widget _buildSuggestions() {
    return new ListView.builder(
        padding: const EdgeInsets.all(16.0),

        // 对于每个建议的单词对都会调用一次 itemBuilder，
        // 然后将单词对添加到 ListTile 行中
        // 在偶数行，该函数会为单词对添加一个 ListTile row.
        // 在奇数行，该函数会添加一个分割线的 widget，来分隔相邻的词对。
        // 注意，在小屏幕上，分割线看起来可能比较吃力。

        itemBuilder: (BuildContext _context, int i) {
          // 在每一列之前，添加一个1像素高的分隔线widget
          if (i.isOdd) {
            return new Divider();
          }

          // 语法 "i ~/ 2" 表示i除以2，但返回值是整形（向下取整）
          // 比如 i 为：1, 2, 3, 4, 5 时，结果为 0, 1, 1, 2, 2，
          // 这可以计算出 ListView 中减去分隔线后的实际单词对数量
          final int index = i ~/ 2;
          // 如果是建议列表中最后一个单词对
          if (index >= _suggestions.length) {
            // ...接着再生成10个单词对，然后添加到建议列表
            _suggestions.addAll(generateWordPairs().take(10));
          }
          return _buildRow(_suggestions[index]);
        });
  }

  @override // 新增代码片段 - 开始 ...
  Widget build(BuildContext context) {
    //final wordPair = new WordPair.random(); // 删掉 ...
    //return new Text(wordPair.asPascalCase); // ... 这两行

    return new Scaffold(
      // 代码从这里...
      appBar: new AppBar(
        title: new Text('Startup Name Generator'),
        actions: <Widget>[
          // 新增代码开始 ...
          new IconButton(icon: const Icon(Icons.list), onPressed: _pushSaved),
        ],
      ),
      body: _buildSuggestions(),
    );
  } // ... 新增的代码片段 - 结束

  // 新增代码开始
  void _pushSaved() {
    Navigator.of(context).push(
      new MaterialPageRoute<void>(
        // 新增如下20行代码 ...
        builder: (BuildContext context) {
          final Iterable<ListTile> tiles = _saved.map(
            (WordPair pair) {
              return new ListTile(
                title: new Text(
                  pair.asPascalCase,
                  style: _biggerFont,
                ),
              );
            },
          );
          final List<Widget> divided = ListTile.divideTiles(
            context: context,
            tiles: tiles,
          ).toList();
          return new Scaffold(
            // 新增 6 行代码开始 ...
            appBar: new AppBar(
              title: const Text('Saved Suggestions'),
            ),
            body: new ListView(children: divided),
          ); // ... 新增代码段结束.
        },
      ), // ... 新增代码结束
    );
  }

  Widget _buildRow(WordPair pair) {
    final bool alreadySaved = _saved.contains(pair); // 新增本行
    return new ListTile(
      title: new Text(
        pair.asPascalCase,
        style: _biggerFont,
      ),
      trailing: new Icon(
        // 新增代码开始 ...
        alreadySaved ? Icons.favorite : Icons.favorite_border,
        color: alreadySaved ? Colors.red : null,
      ), // ... 新增代码结束
      onTap: () {
        // 增加如下 9 行代码...
        setState(() {
          if (alreadySaved) {
            _saved.remove(pair);
          } else {
            _saved.add(pair);
          }
        });
      }, // ... 一直到这里
    );
  }
}

Future<List<Photo>> fetchPhotos(http.Client client) async {
  final response =
      await client.get('https://jsonplaceholder.typicode.com/photos');

  // Use the compute function to run parsePhotos in a separate isolate.
  return compute(parsePhotos, response.body);
}

// A function that converts a response body into a List<Photo>.
List<Photo> parsePhotos(String responseBody) {
  final parsed = jsonDecode(responseBody).cast<Map<String, dynamic>>();

  return parsed.map<Photo>((json) => Photo.fromJson(json)).toList();
}

class Photo {
  final int albumId;
  final int id;
  final String title;
  final String url;
  final String thumbnailUrl;

  Photo({this.albumId, this.id, this.title, this.url, this.thumbnailUrl});

  factory Photo.fromJson(Map<String, dynamic> json) {
    return Photo(
      albumId: json['albumId'] as int,
      id: json['id'] as int,
      title: json['title'] as String,
      url: json['url'] as String,
      thumbnailUrl: json['thumbnailUrl'] as String,
    );
  }
}

class MyHomePage extends StatelessWidget {
  final String title;
	// Follow this tutorial https://medium.com/flutter-community/using-svg-in-flutter-3dcf7b1dd713
	final String dogUrl = 'https://www.svgrepo.com/show/2046/dog.svg';
	final String dogFoodUrl = 'https://www.svgrepo.com/show/3682/dog-food.svg';

  MyHomePage({Key key, this.title}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: Center(
        child: FloatingActionButton(
					onPressed: () {
						print('Thanks');
					},
					child: SvgPicture.network(
						dogFoodUrl,
						semanticsLabel: 'Feed button',
						placeholderBuilder: (context) => Icon(Icons.error),
					),
				),
//        FutureBuilder<List<Photo>>(
//        future: fetchPhotos(http.Client()),
//        builder: (context, snapshot) {
//          if (snapshot.hasError) print(snapshot.error);
//
//          return snapshot.hasData
//              ? PhotosList(photos: snapshot.data)
//              : Center(child: CircularProgressIndicator());
//        }),
      ),
    );
  }
}

class PhotosList extends StatelessWidget {
  final List<Photo> photos;

  PhotosList({Key key, this.photos}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GridView.builder(
      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 2,
      ),
      itemCount: photos.length,
      itemBuilder: (context, index) {
        return Image.network(photos[index].thumbnailUrl);
      },
    );
  }
}
