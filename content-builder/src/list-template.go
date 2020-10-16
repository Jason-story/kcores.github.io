// list-template.go
package main

const LIST_P1 = `<!DOCTYPE html>
<html>

<head>
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-180293201-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-180293201-1');
</script>

<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>  KCORES - 氪金核心 </title>
<link rel="shortcut icon" href="assets/logo/kcores-logo.ico"/>
<meta name="keywords" content=" KCORES, 氪金核心, 消费电子, 万兆网络, 家用NAS, 垃圾佬, 服务器, 云服务及主机, 仪表, Homelab, Vintage" />
<meta name="description" content="KCORES - 氪金核心. 一个奇特的电子产品爱好者网站. 该项目由 @karminski-牙医 发起, 目的是建立一个大家能轻松讨论泛计算机话题的环境." />
<meta name="viewport" content="width=device-width, user-scalable=no">
<!-- include scripts & styles -->
<link rel="stylesheet" type="text/css" href="assets/styles/reset.css">
<link rel="stylesheet" type="text/css" href="assets/styles/base.css">
<link rel="stylesheet" type="text/css" href="assets/styles/list.css">
<script src="assets/js/jquery.min.js"></script>
<style type="text/css">
    .water-basic {
      position: relative;
      margin: 2em auto 0 auto; 
      width: 1008px;
    }
    .item {
      position: absolute;
      width: 236px;
      margin: 5px;
      transition: all 1s;
      border-radius: 1em;
      margin-bottom: 2em;
    }
    .entry-image {
      width: 236px;
      border-radius: 1em;
    }
    .entry-title {
      padding: 0.4em 0em 0 0.4em;
      font-weight: 600;
      line-height: 1.2em;
      font-size: 1em;
    }
    .entry-info {
      padding: 0.4em 0.4em 0 0.4em;
      font-weight: 400;
      font-size: 0.6em;
      color: #909090;
    }
`

const LIST_P2 = `
	.box%d {
      height: %dpx;
      background-image: url(%s);
    }
`

const LIST_P3 = `
   
  </style>

</head>

<body>

<div id="main">

    <!-- Herader -->
    <header>
      <div class="inner cf">
        <nav class="left">
          <a href="https://kcores.com/" class="logo"><h1>KCORES 氪金核心</h1></a>
          <a href="https://kcores.com/reading" class="nav-button"><h1>阅读</h1></a>
          <a href="https://kcores.com/topics" class="nav-button"><h1>话题</h1></a>
        </nav>
        <nav class="right">
          <a href="https://kcores.com/about">关于</a>
          <a href="https://kcores.com/blogroll">友情链接</a>
          <a target="_blank" href="https://github.com/KCORES">Github</a>
      </nav></div>
    </header>

    <!-- masthead -->
    <div id="masthead">
      <div id="masthead-content">
        <div id="masthead-icon">
            <img src="%s" />
        </div>
        <div id="masthead-subcontent">
          <h1 id="masthead-title">%s</h1>
          <p id="masthead-desc">%s</p>
        </p>
      </div>
    </div>

</div>




<!-- grid content -->
<div class="water-basic">
`

const LIST_P4 = `
  <!-- THIS CONTENT IS AUTOMATIC GENERATED BY CONTENT-BUILDER, DO NOT EDIT THIS FILE MANUALLY -->
  <div class="item">
    <div class="entry-image box%d"></div>
    <div class="entry-title"><a target="_blank" href="%s">%s</a></div>
    <div class="entry-info">%s / %s</div>
  </div>
`

const LIST_P5 = `
</div>

<!-- load js -->
<script src="assets/js/list-page.js"></script>


</body>
</html>
`