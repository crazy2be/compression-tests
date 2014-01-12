var zip = require("./inflate");
var arrays = require("arrays.js");

process.stdin.resume();
process.stdin.setEncoding('utf8');

var inflater = new zip.Inflater();

process.stdin.on('data', function(chunk) {
  var buf = new Buffer(chunk, 'base64');
  var arr = new arrays.Uint8Array(buf);
  var data = inflater.append(arr, function(){});
  process.stdout.write(decodeUTF8(data));
});

process.stdin.on('end', function () {
  inflater.flush();
});

// Not really UTF-8 but close enough.
function decodeUTF8(arr) {
  var str = '';
  for (var i = 0; i < arr.length; i++) {
    var c = arr[i];
    str = str + String.fromCharCode(c);
  }
  return str;
}