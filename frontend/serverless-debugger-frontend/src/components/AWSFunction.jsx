import React, { useState, useEffect } from "react";
import {fetchAWSFunctions, invokeFunction, updateFunction} from "../api";

const AWSFunctions = () => {
    const [functions, setFunctions] = useState([]);

    useEffect (() =>{
        async function loadFunctions(){
            const data = await fetchAWSFunctions();
            setFunctions(data);
        }
        loadFunctions();
    },[]);

    const handleInvoke = (functionName) => {
        invokeFunction("aws", functionName).then(response => {
            console.log(response);
        });
    };

    const handleUpdate = (functionName) => {
        const updateData = {}; //Collect update data from a form of input fields

        updateFunction("aws", functionName, updateData)
            .then(response => {
                console.log(response);
            });
    };

    return (
        <div>
            <h2>AWS Functions</h2>
            <ul>
                {functions.map(fn => (
                    <li key = {fn.name}>
                        {fn.name}
                        <button onClick={() => handleInvoke(fn.name)}>Invoke</button>
                        <button onClick={() => handleUpdate(fn.name)}>Update</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default AWSFunctions;