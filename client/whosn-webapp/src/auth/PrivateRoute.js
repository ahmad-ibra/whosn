import { Navigate } from "react-router-dom";
import jwt from 'jwt-decode'
import { useLocation, useNavigate } from "react-router-dom"

const PrivateRoute = ({ children }) => {

    const destPath = useLocation().pathname
    const navigate = useNavigate();


    console.log(`destPath is ${destPath}`)
    const jwtToken = localStorage.getItem('jwt')
    if (jwtToken == null) {
        return <Navigate to='/login' />
    }

    const decodedToken = jwt(jwtToken);

    // seconds elapsed since January 1, 1970 UTC
    const now = Math.round((new Date()).getTime() / 1000);

    // check not expired (we consider it expired 5 minutes early)
    if (decodedToken.exp - 300 < now) {
        console.log("expired")
        localStorage.removeItem('jwt')
        return navigate('/login', { destPath });
    }

    // check issuer
    if (decodedToken.iss !== "whosn-core") {
        console.log("wrong issuer")
        localStorage.removeItem('jwt')
        return navigate('/login', { destPath });
    }

    return children
}

export default PrivateRoute
