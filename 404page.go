package main

const custom404Content = `
<html>
	<head>
		<title>not found</title>
		<style type="text/css">
body {
  text-align: center;
  background-color: #1d1f20;
}

.wrapper {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 90%;
  font-size: 0;
  -webkit-transform: translate(-50%);
          transform: translate(-50%);
}

p {
  font-family: sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #eee;
  opacity: 0.3;
}

.letter {
  width: 24px;
  display: inline-block;
  vertical-align: middle;
  position: relative;
  overflow: hidden;
  margin: 0 0;
  font-family: sans-serif;
  font-size: 24px;
  font-weight: 600;
  line-height: 24px;
  text-transform: uppercase;
  color: #eee;
}
.letter:before {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  word-break: break-all;
  background-color: #1d1f20;
}

.letter:nth-child(1):before {
  content: "YV3EI2WA7SFP6DXR415N0MCGL9UOHZTBK8Q";
  margin-top: -792px;
  -webkit-animation-name: letter1;
          animation-name: letter1;
  -webkit-animation-duration: 0.4076470588s;
          animation-duration: 0.4076470588s;
  -webkit-animation-delay: 0.94s;
          animation-delay: 0.94s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter1 {
  from {
    margin-top: -792px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter1 {
  from {
    margin-top: -792px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(2):before {
  content: "DXIZQOR76SNGME9W3LAV2TYF0B8415CHPKU";
  margin-top: -216px;
  -webkit-animation-name: letter2;
          animation-name: letter2;
  -webkit-animation-duration: 4.725s;
          animation-duration: 4.725s;
  -webkit-animation-delay: 0.25s;
          animation-delay: 0.25s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter2 {
  from {
    margin-top: -216px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter2 {
  from {
    margin-top: -216px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(3):before {
  content: "3Y0LRM5ZQUBPI7D19FH68NVKEO2XTAC4GSW";
  margin-top: -624px;
  -webkit-animation-name: letter3;
          animation-name: letter3;
  -webkit-animation-duration: 1.6851851852s;
          animation-duration: 1.6851851852s;
  -webkit-animation-delay: 0.75s;
          animation-delay: 0.75s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter3 {
  from {
    margin-top: -624px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter3 {
  from {
    margin-top: -624px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(4):before {
  content: "1XSGRB5TO7LHCNZIYEK8WDP42UQ3AMF690V";
  margin-top: -384px;
  -webkit-animation-name: letter4;
          animation-name: letter4;
  -webkit-animation-duration: 1.7129411765s;
          animation-duration: 1.7129411765s;
  -webkit-animation-delay: 0.74s;
          animation-delay: 0.74s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter4 {
  from {
    margin-top: -384px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter4 {
  from {
    margin-top: -384px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(5):before {
  content: "V0XA2KGP8Q5FI1C764ZLBUYHD9SRTNEW3MO";
  margin-top: 0px;
  -webkit-animation-name: letter5;
          animation-name: letter5;
  -webkit-animation-duration: 0s;
          animation-duration: 0s;
  -webkit-animation-delay: 0.91s;
          animation-delay: 0.91s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter5 {
  from {
    margin-top: 0px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter5 {
  from {
    margin-top: 0px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(6):before {
  content: "07BPA2SQH9RWVT8Y3EFLMCGUOK45ZNXI6D1";
  margin-top: -504px;
  -webkit-animation-name: letter6;
          animation-name: letter6;
  -webkit-animation-duration: 2.8063636364s;
          animation-duration: 2.8063636364s;
  -webkit-animation-delay: 0.58s;
          animation-delay: 0.58s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter6 {
  from {
    margin-top: -504px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter6 {
  from {
    margin-top: -504px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(7):before {
  content: "DQ2V639Z4AUIRX1TLFYP8KBEWH75N0COSMG";
  margin-top: -120px;
  -webkit-animation-name: letter7;
          animation-name: letter7;
  -webkit-animation-duration: 4.55s;
          animation-duration: 4.55s;
  -webkit-animation-delay: 0.22s;
          animation-delay: 0.22s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter7 {
  from {
    margin-top: -120px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter7 {
  from {
    margin-top: -120px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(8):before {
  content: "Y1LFRBGUZHT95476SIAXE3PC0MOQDVKN8W2";
  margin-top: -744px;
  -webkit-animation-name: letter8;
          animation-name: letter8;
  -webkit-animation-duration: 1.085s;
          animation-duration: 1.085s;
  -webkit-animation-delay: 0.84s;
          animation-delay: 0.84s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter8 {
  from {
    margin-top: -744px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter8 {
  from {
    margin-top: -744px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(9):before {
  content: "UXGTZWMKV8QD07YSP3A62INF95BLHECR1O4";
  margin-top: -216px;
  -webkit-animation-name: letter9;
          animation-name: letter9;
  -webkit-animation-duration: 2.898s;
          animation-duration: 2.898s;
  -webkit-animation-delay: 0.54s;
          animation-delay: 0.54s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter9 {
  from {
    margin-top: -216px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter9 {
  from {
    margin-top: -216px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(10):before {
  content: "APVI35EGUXYRO6972MZCFSB0T8W4N1HQKDL";
  margin-top: -336px;
  -webkit-animation-name: letter10;
          animation-name: letter10;
  -webkit-animation-duration: 4.1813333333s;
          animation-duration: 4.1813333333s;
  -webkit-animation-delay: 0.36s;
          animation-delay: 0.36s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter10 {
  from {
    margin-top: -336px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter10 {
  from {
    margin-top: -336px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(11):before {
  content: "4ADEOX9LSCZUW1FKIRB6VY8730M52GQTPNH";
  margin-top: -384px;
  -webkit-animation-name: letter11;
          animation-name: letter11;
  -webkit-animation-duration: 2.0423529412s;
          animation-duration: 2.0423529412s;
  -webkit-animation-delay: 0.69s;
          animation-delay: 0.69s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter11 {
  from {
    margin-top: -384px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter11 {
  from {
    margin-top: -384px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(12):before {
  content: "96RCP5BQ30E2SNA14XKTOMWDFUV8LG7YHZI";
  margin-top: -360px;
  -webkit-animation-name: letter12;
          animation-name: letter12;
  -webkit-animation-duration: 3.346875s;
          animation-duration: 3.346875s;
  -webkit-animation-delay: 0.49s;
          animation-delay: 0.49s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter12 {
  from {
    margin-top: -360px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter12 {
  from {
    margin-top: -360px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(13):before {
  content: "EG0AU614M9H7DI8SFTYOXNKRQ5W3VZL2BPC";
  margin-top: -720px;
  -webkit-animation-name: letter13;
          animation-name: letter13;
  -webkit-animation-duration: 5.3516129032s;
          animation-duration: 5.3516129032s;
  -webkit-animation-delay: 0.21s;
          animation-delay: 0.21s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter13 {
  from {
    margin-top: -720px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter13 {
  from {
    margin-top: -720px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(14):before {
  content: "CU4IBG1KRV06F87ZL9XNMWED5AOSQT23PYH";
  margin-top: -816px;
  -webkit-animation-name: letter14;
          animation-name: letter14;
  -webkit-animation-duration: 3.74s;
          animation-duration: 3.74s;
  -webkit-animation-delay: 0.45s;
          animation-delay: 0.45s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter14 {
  from {
    margin-top: -816px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter14 {
  from {
    margin-top: -816px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(15):before {
  content: "E9TXAC1SWH7UY2PZ64G0O5NVF3MQKRLIBD8";
  margin-top: -264px;
  -webkit-animation-name: letter15;
          animation-name: letter15;
  -webkit-animation-duration: 1.155s;
          animation-duration: 1.155s;
  -webkit-animation-delay: 0.82s;
          animation-delay: 0.82s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter15 {
  from {
    margin-top: -264px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter15 {
  from {
    margin-top: -264px;
  }
  to {
    margin-top: 24px;
  }
}
.letter:nth-child(16):before {
  content: "UAH0CZS3V24M7L6IOW8QE5GBRTY1DPKXNF9";
  margin-top: -720px;
  -webkit-animation-name: letter16;
          animation-name: letter16;
  -webkit-animation-duration: 6.3s;
          animation-duration: 6.3s;
  -webkit-animation-delay: 0.07s;
          animation-delay: 0.07s;
  -webkit-animation-fill-mode: forwards;
          animation-fill-mode: forwards;
}

@-webkit-keyframes letter16 {
  from {
    margin-top: -720px;
  }
  to {
    margin-top: 24px;
  }
}

@keyframes letter16 {
  from {
    margin-top: -720px;
  }
  to {
    margin-top: 24px;
  }
}

		</style>
	</head>
	<body>
		<div class="wrapper">
			<div class="letters">
				<span class="letter">4</span>
				<span class="letter">0</span>
				<span class="letter">4</span>
				<span class="letter">&nbsp;</span>
				<span class="letter">n</span>
				<span class="letter">o</span>
				<span class="letter">t</span>
				<span class="letter">&nbsp;</span>
				<span class="letter">f</span>
				<span class="letter">o</span>
				<span class="letter">u</span>
				<span class="letter">n</span>
				<span class="letter">d</span>
			</div>
			<p>http-fs</p>
		</div>
	</body>
</html>
`
