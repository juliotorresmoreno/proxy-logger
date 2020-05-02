<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Logger</title>
</head>

<body>
  {{range $index, $request := .}}
  <div style="cursor: pointer;" onclick='javascript: show({{$index}})'>
    {{$request.Time}} {{$request.Method}} {{$request.URI}} {{$request.StatusCode}}
  </div>
  <div id='content_{{$index}}' class='request'>
    <pre>{{$request.RawRequest}}</pre>
    <pre>{{$request.RawResponse}}</pre>
  </div>
  {{ end }}
  <style type='text/css'>
    .request {
      display: none
    }
  </style>
  <script type='text/javascript'>
    function show(id) {
      document.querySelectorAll('.request')
        .forEach(function (el) {
          el.style.display = 'none'
        })
      let el = document.getElementById(`content_${id}`)
      if (el.getAttribute('visible') === 'true') {
        el.setAttribute('visible', 'false')
        el.style.display = 'none';
      } else {
        el.setAttribute('visible', 'true')
        el.style.display = 'block';
      }
    }
  </script>
</body>

</html>