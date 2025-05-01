import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Main from './pages/Main';
import Login from './pages/Login';
import Sidebar from './components/Sidebar/Sidebar';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route
          path="/main"
          element={
            <>
              <Sidebar />
              <Main />
            </>
          }
        />
      </Routes>
    </Router>
  );
};

export default App;