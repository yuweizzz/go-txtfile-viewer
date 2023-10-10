<html>
  <head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no"/>
    <style>
      html body {
        font-size: 16px;
      }
      #main {
        margin: 0 auto;
        width: 700px;
        word-wrap: break-word;
      }
      #content {
        padding: 15px;
      }
      p {
        margin-bottom: 16px;
      }
      @media screen and (max-width: 700px) {
        #main {
          width: 100%;
          font-size: 18px;
          word-wrap: break-word;
        }
      }
      * {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <div id="main"><div id="content">{{ .Content }}</div></div>
  </body>
</html>
