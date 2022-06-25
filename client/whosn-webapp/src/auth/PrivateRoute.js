import { Navigate } from "react-router-dom";
import jwt from 'jwt-decode'

const PrivateRoute = ({ children }) => {

    const jwtToken = localStorage.getItem('jwt')
    if (jwtToken == null) {
        return <Navigate to='/login' />
    }

    const decodedToken = jwt(jwtToken);

    // seconds elapsed since January 1, 1970 UTC
    const now = Math.round((new Date()).getTime() / 1000);

    // check not expired (we consider it expired 5 minutes early)
    if (decodedToken.exp - 300 < now) {
        localStorage.removeItem('jwt')
        return <Navigate to='/login' />
    }

    // check issuer
    if (decodedToken.iss !== "whosn-core") {
        localStorage.removeItem('jwt')
        return <Navigate to='/login' />
    }

    return children
}

export default PrivateRoute
