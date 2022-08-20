local uri = ngx.var.uri;
local path = "%s"
local regex = "%s"

if (path ~= "" and uri == path) or (regex ~= "" and string.match(uri, regex))
then
    ngx.header["a"] = "b"
    ngx.header["Content-Type"] = "text/plain"
    ngx.say(uri);
    ngx.exit(200);
end
