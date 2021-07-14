let ws;

// ページから離脱時に発生
window.onbeforeunload = () => {
  console.log('User Leaving.');
  let jsonData = {};
  jsonData['action'] = 'left';
  ws.send(JSON.stringify(jsonData));
}

function connect() {
  let loc = window.location;
  let uri = 'ws:';

  if (loc.protocol === 'https:') {
    uri = 'wss:';
  }

  uri += '//' + loc.host;
  uri += loc.pathname + 'ws';
  ws = new WebSocket(uri);

  ws.onopen = () => {
    console.log('Connected.');
  }

  ws.onmessage = e => {
    let data = JSON.parse(e.data);
    console.log({data});
    output.innerHTML += data.message + '<br>';
  }

  ws.onclose = e => {
    console.log('connection closed. 3秒後に再接続します。');
    setTimeout(() => {
      connect();
    }, 3000);
  }

  ws.onerror = e => {
    console.log('there was an error.');
    ws.close();
  }
}

document.addEventListener('DOMContentLoaded', () => {
  const input = document.getElementById('input');
  const output = document.getElementById('output');
  const btn = document.querySelector('.btn');
  connect();

  btn.addEventListener('click', () => {
    let jsonData = {};
    jsonData['action'] = 'broadcast';
    jsonData['message'] = input.value;
    ws.send(JSON.stringify(jsonData));
    input.value = '';
  })
});
