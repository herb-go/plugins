function test(url){
    let req= HTTP.NewRequest("GET",url)
    req.GetID()
    req.GetURL()
    req.SetURL(url)
    req.GetBody()
    req.SetBody("test")
    req.ResetHeader()
    req.SetHeader("uid","123")
    req.AddHeader("uid","234")
    req.GetHeader("uid")
    req.DelHeader("uid")
    req.GetMethod()
    req.SetMethod("POST")
    req.HeaderValues("uid")
    req.HeaderFields()
    req.Execute()
    req.ExecuteStatus()
    req.FinishedAt()
    req.ResponseStatusCode()
    req.ResponseStatusCode()
    req.ResponseBody()
    req.ResponseHeader()
    req.ResponseHeaderValues()
    req.ResponseHeaderFields()
}