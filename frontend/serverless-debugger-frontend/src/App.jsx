import { useState } from 'react'
import './App.css'
import AWSFunctions from './components/AWSFunctions';
import GCPFunctions from './components/GCPFunctions';
import FunctionLogs from './components/FunctionLogs';
import Debugger from './components/Debugger';


function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
            <h1>Serverless Debugger</h1>
            <AWSFunctions />
            <GCPFunctions />
            <FunctionLogs />
            <Debugger />
    </div>
  )
}

export default App
