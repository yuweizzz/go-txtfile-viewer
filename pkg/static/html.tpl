<html>
  <head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no"/>
    <style>
      * {
        margin: 0;
        padding: 0;
      }
      html {
        font-size: 16px;
      }
      p,h1,h2,h3,h4,h5,h6 {
        margin-bottom: 16px;
      }
      ol, ul {
        margin: 0;
        padding: 0;
        list-style-type: none;
      }
      ol + p {
        margin-top: 16px;
      }
      ul + p {
        margin-top: 16px;
      }
      ol {
        list-style-type: decimal;
        margin-left: 20px;
      }
      ul {
        list-style-type: disc;
        margin-left: 20px;
      }
      blockquote {
        display: block;
        padding: 0 1em;
        color: #6a737d;
        border-left: 0.25em solid #dfe2e5;
      }
      pre {
        display: block;
        overflow: auto;
        background-color: #f6f8fa;
        padding: 16px;
        border-radius: 6px;
        margin-bottom: 16px;
        font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
        font-size: 90%;
      }
      code {
        display: inline;
        overflow: visible;
        background-color: #f6f8fa;
        word-wrap: normal;
        line-height: inherit;
        max-width: auto;
        border: 0;
        padding: 0;
        margin: 0;
        font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
        font-size: 90%;
      }
      #main {
        margin: 0 auto;
        width: 700px;
        word-wrap: break-word;
      }
      #content {
        padding: 15px;
      }
      @media screen and (max-width: 700px) {
        #main {
          width: 100%;
          font-size: 18px;
          word-wrap: break-word;
        }
      }
    </style>
  </head>
  <body>
    <div id="main"><div id="content">{{ .Content }}</div></div>
  </body>
</html>
