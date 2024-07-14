import React, { useState } from 'react';
import { addBreakPoint } from '../api';

const Debugger = () => {
    const [service, setService] = useState('');
    const [functionName, setFunctionName] = useState('');
    const [fileName, setFileName] = useState('');
    const [lineNumber, setLineNumber] = useState('');

    const handleAddBreakPoint = () => {
        addBreakPoint(service, functionName, fileName, lineNumber).then(response => {
            console.log(response);
        });
    };

    return (
        <div>
            <h2>Debugger</h2>
            <select onChange={(e) => setService(e.target.value)}>
                <option value="">Select Service</option>
                <option value="aws">AWS</option>
                <option value="gcp">GCP</option>
            </select>
            <input
                type="text"
                placeholder="Function Name"
                value={functionName}
                onChange={(e) => setFunctionName(e.target.value)}
            />
            <input
                type="text"
                placeholder="File Name"
                value={fileName}
                onChange={(e) => setFileName(e.target.value)}
            />
            <input
                type="text"
                placeholder="Line Number"
                value={lineNumber}
                onChange={(e) => setLineNumber(e.target.value)}
            />
            <button onClick={handleAddBreakPoint}>Add Breakpoint</button>
        </div>
    );
};

export default Debugger;
