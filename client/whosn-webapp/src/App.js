import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './routes/Login';
import Home from './routes/Home';
import Register from './routes/Register';
import DetailedEvent from './routes/DetailedEvent';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path='/' element={<Home />} />
        <Route exact path='/login' element={<Login />} />
        <Route exact path='/register' element={<Register />} />
        <Route exact path="/event/*" element={<DetailedEvent />} />
      </Routes>
    </Router>
  );
}

export default App;
