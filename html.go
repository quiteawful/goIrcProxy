package main

var HtmlMain string = `<!DOCTYPE html>
<html>
<head>
    <meta charset=utf-8 />
    <title>Doclol Irc Proxy</title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        .wrap {
            max-width: 320px;
        }

        .content {
            height: 300px;
        }

        * {
            box-sizing: border-box;
            font-size: 12px;
            color: #aaaaaa;
        }

        .wrap {
            width: 100%;
            padding-bottom: 30px;
            position: relative;
        }

        .form {
            position: relative;
            width: 100%;
            height: 30px;
            margin: 0;
            padding: 0;
            border: 1px solid #eee;
            border-top: 0;
            padding-right: 70px;
        }

        .input {
            border: 0;
            margin: 0;
            padding: 0px 10px;
            width: 100%;
            display: block;
            line-height: 30px;
            background: transparent;
        }

        .submit {
            float: right;
            position: absolute;
            top: 0;
            right: 0;
            width: 70px;
            border: 0;
            border-left: 1px solid #eee;
            color: #eee;
            background-color: transparent;
            margin: 0;
            padding: 0;
            display: block;
            height: 100%;
            line-height: 30px;
        }

        .content {
            width: 100%;
            overflow-y: scroll;
            border: 1px solid #eee;
            padding: 10px;
            margin: 0;
            display: block;
            word-wrap: break-word;
            white-space: pre-wrap;
            white-space: -moz-pre-wrap;
            white-space: -pre-wrap;
            white-space: -o-pre-wrap;
        }
    </style>
</head>
<body>
    <div class="wrap">
        <pre id="content" class="content">

        </pre>
        <form id="form" class="form" action="/" method="post">
            <input class="input" type="text" name="content" placeholder="message...">
            <input class="submit" type="submit" name="text" value="send">
        </form>
    </div>

    <script type="text/javascript">
        var refreshInterval = 5000;
        var $form = document.querySelector('#form');
        var $input = $form.querySelector('.input');

        // Event bindings
        addEvent($form, 'submit', onFormSubmit);

        // Initial request
        updateLog();

        // Setup request Interval
        window.setInterval(function(){
            updateLog();
        }, refreshInterval);

        $input.focus();

        // Request log and replace content
        function updateLog(){
            xhr('/log', function(response){
                var $content = document.querySelector('#content');
                $content.innerHTML = response;
                $content.scrollTop = $content.scrollHeight;
            });
        }

        // Send form content, reset input and update log
        function onFormSubmit(event){
            event.preventDefault();
            var $input = $form.querySelector('.input');
            var content = $input.value;

            if(content.length == 0){
                return;
            }

            var data = 'content=' + content;
            xhr('/', function(){
                $input.value = '';
                updateLog();
            }, 'POST', data);
        }

        // XHR stuff
        function xhr(url, callback, method, data) {
            method = typeof method !== 'undefined' ? method : 'GET';
            var xhReq;

            if (window.XMLHttpRequest) {
                xhReq = new XMLHttpRequest();
            } else {
                xhReq = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xhReq.onreadystatechange = function() {
                if (xhReq.readyState == 4 ) {
                    if(xhReq.status == 200){
                        callback(xhReq.responseText);
                    }
                }
            }

            if (method == 'GET'){
                if(typeof data !== 'undefined' && data !== null) {
                    url = url + '?' + data;
                }
                data = null;
            }

            xhReq.open(method, url, true);
            xhReq.setRequestHeader("X-Requested-With", "xhReqRequest");

            if (method == 'POST'){
                xhReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
            }

            xhReq.send(data);
        }

        // Event stuff
        function addEvent(el, eventType, handler) {
            if (el.addEventListener) {
                el.addEventListener(eventType, handler, false);
            } else if (el.attachEvent) {
                el.attachEvent('on' + eventType, handler);
            } else {
                el['on' + eventType] = handler;
            }
        }
    </script>
</body>
</html>
`
