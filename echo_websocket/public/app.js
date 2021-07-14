document.addEventListener('DOMContentLoaded', () => {
  // let loc = window.location;
  // let uri = 'ws:';
  //
  // if (loc.protocol === 'https:') {
  //   uri = 'wss:';
  // }
  //
  // uri += '//' + loc.host;
  // uri += loc.pathname + 'ws';
  let uri = 'ws://localhost:8080/ws'
  const input = document.getElementById('input');
  const output = document.getElementById('output');
  const btn = document.querySelector('.btn');
  const ws = new WebSocket(uri);

  ws.onopen = () => {
    console.log('Connected.');
  }

  ws.onmessage = msg => {
    let data = JSON.parse(msg.data);
    output.innerHTML += data.text + '<br>';
  }

  ws.onclose = () => {
    console.log('connection closed.');
  }

  ws.onerror = e => {
    console.log('there was an error.');
  }

  // ページから離脱時に発生
  window.onbeforeunload = function() {
    console.log('User Leaving.');
    let jsonData = {};
    jsonData['action'] = 'left';
    ws.send(JSON.stringify(jsonData));
  }

  btn.addEventListener('click', () => {
    let jsonData = {};
    jsonData['action'] = 'send';
    jsonData['text'] = input.value;
    ws.send(JSON.stringify(jsonData));
    input.value = '';
  })
});
