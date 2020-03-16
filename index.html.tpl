<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Hello from Kubernetes!</title>
  <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet">
</head>
<body>
  <div class="h-full w-full">
    <div class="flex flex-col items-center mt-20">
      <div class="border rounded-lg p-20 shadow-2xl w-1/2">
        Following pods have been found:

        <table class="table-auto mt-5">
          {{ range . }}
            <tr>
                <td class="border p-5">{{ .ObjectMeta.Name }}</td>
                <td class="border p-5">{{ .Status.Message }}</td>
            </tr>
          {{ end }}
        </table>
      </div>
    </div>
  </div>
</body>
</html>
