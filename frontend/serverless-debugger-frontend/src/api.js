const API_URL_BASE = "http://localhost:8080";

export async function fetchAWSFunctions(){
    const response = await fetch(`${API_URL_BASE}/aws/functions`);
    return response.json();
}

export async function fetchGCPFunctions(){
    const response = await fetch(`${API_URL_BASE}/gcp/functions`);
    return response.json();
}

export async function invokeFunction(service, functionName){
    const response = await fetch(`${API_URL_BASE}/${service}/invoke/${functionName}`,{
        method: 'POST'
    })
    return response.json();

}

export async function updateFunction(service, functionName, updateData){
    const response = await fetch(`${API_URL_BASE}/${service}/update/${functionName}`,{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(updateData)
    })
    return response.json();
}

export async function fetchLogs(service, functionName, startTime, endTime){
    const response = await fetch(`${API_URL_BASE}/${service}/logs/${functionName}?startTime=${startTime}&endTime=${endTime}`);
    return response.json();
}

export async function addBreakPoint(service, functionName, fileName, lineNumber){
    const response = await fetch(`${API_URL_BASE}/${service}/debugger/addBreakPoint/${functionName}?fileName=${fileName}&lineNumber=${lineNumber}`,{
        method: 'POST'
    })
    return response.json();
}
