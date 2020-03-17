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
        <div>
          Following pods have been found:
        </div>

        <div class="mt-8">
          <table class="table-auto">
            {{ range . }}
              <tr>
                  <td class="border p-5">{{ .Name }}</td>
                  <td class="border p-5">{{ .Status.Phase }}</td>
              </tr>
            {{ end }}
          </table>
        </div>

        <div class="mt-16">
          <a class="float-left hover:text-yellow-600" href="https://gitlab.com/grzesiek/hello-pods">
            See <i>grzesiek/hello-pods</i> on &nbsp;
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#fca326" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="float-right feather feather-gitlab"><path d="M22.65 14.39L12 22.13 1.35 14.39a.84.84 0 0 1-.3-.94l1.22-3.78 2.44-7.51A.42.42 0 0 1 4.82 2a.43.43 0 0 1 .58 0 .42.42 0 0 1 .11.18l2.44 7.49h8.1l2.44-7.51A.42.42 0 0 1 18.6 2a.43.43 0 0 1 .58 0 .42.42 0 0 1 .11.18l2.44 7.51L23 13.45a.84.84 0 0 1-.35.94z"></path></svg>
          </a>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
