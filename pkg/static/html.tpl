<html>
<head>
    <title>{{ .Title }}</title>
    <style>
        #parent {
            position: absolute;
            width: 100%;
            height: 100%;
            display: flex;
            justify-content: center;
        }
        #left {
            flex: auto;
        }
        #center {
            width: 700px;
        }
        #right {
            flex: auto;
        }
        @media screen and (max-width: 700px) {
            #left {
                display: none;
            }
            #center {
                width: 100%;
            }
            #right {
                display: none;
            }
        }
        * {
            margin: 0;
            padding: 0;
        }
    </style>
</head>
<body>
<div id="parent">
    <div id="left"></div>
    <div id="center">{{ .Content }}</div>
    <div id="right"></div>
</div>
</body>
</html>
