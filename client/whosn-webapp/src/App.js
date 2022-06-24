import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Login from './routes/Login';
import Home from './routes/Home';
import Register from './routes/Register';
import DetailedEvent from './routes/DetailedEvent';
import NotFound from './routes/NotFound';
import PrivateRoute from './auth/PrivateRoute';

function App() {
  return (
    < Router >
      <div className={'main-container'}>
        <Routes>
          <Route exact path='/login' element={<Login />} />
          <Route exact path='/register' element={<Register />} />
          <Route exact path='/' element={
            <PrivateRoute>
              <Home />
            </PrivateRoute>} />
          <Route exact path='/event/*' element={
            <PrivateRoute>
              <DetailedEvent />
            </PrivateRoute>} />
          <Route path='*' element={<NotFound />} />
        </Routes>
      </div>
    </Router >
  );
}

export default App;
