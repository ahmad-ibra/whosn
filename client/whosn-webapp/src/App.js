import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Login from './routes/Login';
import Home from './routes/Home';
import Register from './routes/Register';
import DetailedEvent from './routes/DetailedEvent';
import NotFound from './routes/NotFound';

function App() {
  return (
    < Router >
      <div className={'main-container'}>
        <Routes>
          <Route exact path='/login' element={<Login />} />
          <Route exact path='/' element={<Home />} />
          <Route exact path='/register' element={<Register />} />
          <Route exact path='/event/*' element={<DetailedEvent />} />
          <Route path='*' element={<NotFound />} />
        </Routes>
      </div>
    </Router >
  );
}

export default App;
