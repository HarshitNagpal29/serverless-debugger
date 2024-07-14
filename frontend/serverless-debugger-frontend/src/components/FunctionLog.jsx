import React, { useState } from "react";
import { fetchLogs } from "../api";

const FunctionLogs = () => {
    const [service, setService] = useState('');
    const [functionName, setFunctionName] = useState('');
    const [startTime, setStartTime] = useState('');
    const [endTime, setEndTime] = useState('');
    const [logs, setLogs] = useState([]);

    const handleFetchLogs = () => {
        fetchLogs(service, functionName, startTime, endTime)
            .then(response => {
                setLogs(response);
            });
    };

    return(
        <div>
            <h2>Function Logs</h2>
            <select onChange={(e) => setService(e.target.value)}>
                <option value="">Select Service</option>
                <option value="aws">AWS</option>
                <option value="gcp">GCP</option>
            </select>
            <input type="text"
             placeholder="Function Name" 
             value={functionName}
             onChange={(e) => setFunctionName(e.target.value)}/>
             <input type="text"
                placeholder="Start Time"
                value={startTime}
                onChange={(e) => setStartTime(e.target.value)}/>
            <input type="text"
                placeholder="End Time"
                value={endTime}
                onChange={(e) => setEndTime(e.target.value)}/>
            <button onClick={handleFetchLogs}>Fetch Logs</button>
            <ul>
                {logs.map((log, index ) => (
                    <li key={index}>
                        {log}
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default FunctionLogs;