
async function call_post(payload: Payload){

    let body = JSON.stringify(payload);

    let res = await fetch("/api/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"     
        },
        body: body 
    })  

    console.log(body)

    gameRedeemer(res)
}
