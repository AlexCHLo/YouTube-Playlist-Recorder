import logo from './logo.svg';
import './App.css';
import Home from './components/Home.js'
import Login from './components/Login.js'
import Overview from './components/Overview.js'
import Topmenu from './components/Topmenu.js'

import {
  BrowserRouter as Router,
  Routes,
  Route} from "react-router-dom";

function App() {
  return (
    <Router>
      <div>
        <Topmenu/>
        <Routes>
          <Route path="/" caseSensitive={false} element={<Home />} />
          <Route path="/login" caseSensitive={false} element={<Login />} />
          <Route path="/overview" caseSensitive={false} element={<Overview />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
