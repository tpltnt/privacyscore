{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <meta name="referrer" content="never" />
    <title>Privacy Score</title>
    <link rel="stylesheet" href="./static/milligram.min.css" />
    <style>
body { background-color: #fcfbfa; margin: 0; padding: 0; }
h2 { font-size: 5.2rem; }
.navigation { border-bottom: 1px solid #ddd; background-color: #f4f5f6; }
.navigation h1 { font-size: 1.4em; padding: 0.4em 0; margin: 0; }
.main { margin-top: 1em; }
#main_form { margin-top: 5em; }
#url_input { color: #606c76; background: #f4f5f6; }
input.big_input { height: 1.6em; padding: 0 0.4em; font-size: 2.2em; line-height: 1.6em; vertical-align: middle; font-weight: 300; }
.invisible { display: none; }
.small { font-size: 0.8em; }

td:first-child, th:first-child { padding-left: 0.6em; }
td:last-child, th:last-child { padding-right: 0.6em; }
th:last-child { text-align: right; }

.result_header { display: flex; align-items: center; }
.result_url { text-align: center; flex-grow: 1; word-wrap: break-word; }
.score { color: #fcfbfa; border-radius: 50%; height: 4em; width: 4em; line-height: 4em; text-align: center; flex-shrink: 0; }
.penalty_link { white-space: nowrap; }
.penalty_desc { font-size: 1.2em; }
.penalty_note { background: #606c76; color: #fcfbfa; font-size: 0.7em; border-radius: 0.4em; padding: 2px 6px; }
.penalty_value { font-size: 1.7em; font-weight: 600; display: block; text-align: right; }
.good { background-color: #40d47e; }
.medium { background-color: #f1c40f; }
.bad { background-color: #e74c3c; }
.poor { background-color: #2c3e50; }
input[type="submit"] { float: right; }
    </style>
</head>
<body>
    <main class="wrapper">
        <nav class="navigation">
            <section class="container">
                <a class="navigation-title" href="./">
                    <h1 class="title">PrivacyScore</h1>
                </a>
            </section>
        </nav>

        <section class="container main">
        {{ template "content" . }}
        </section>
    </main>
</body>
</html>
{{ end }}
