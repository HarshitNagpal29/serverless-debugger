import React, {useEffect, useState} from "react";
import { fetchGCPFunctions, invokeFunction, updateFunction } from "../api";

const GCPFunctions = () => {
    const [functions, setFunctions] = useState([]);

    useEffect(() => {
        async function loadFunctions(){
            const data = await fetchGCPFunctions();
            setFunctions(data);
        }
        loadFunctions();

    }, []);
    const handleInvoke = (functionName) => {
        invokeFunction("gcp", functionName)
            .then(response => {
                console.log(response);
            });
    };
    const handleUpdate = (functionName) => {
        const updateData = {}; //Collect update data from a form of input fields

        updateFunction("gcp", functionName, updateData)
            .then(response => {
                console.log(response);
            });
    };

    return (
        <div>
            <h2>GCP Functions</h2>
            <ul>
                {functions.map(fn => (
                    <li key={fn.name}>
                        {fn.name}
                        <button onClick={() => handleInvoke(fn.name)}>Invoke</button>
                        <button onClick={() => handleUpdate(fn.name)}>Update</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default GCPFunctions;