import { useState, useEffect } from 'react'
import { Navigate, useLocation } from 'react-router-dom'
import jwt from 'jwt-decode'

const PrivateRoute = ({ children }) => {
    const location = useLocation()

    const [path, setPath] = useState('/')
    useEffect(() => {
        const path = `/login?redirectTo=${location.pathname}`
        setPath(path)
    }, [location.pathname])

    const jwtToken = localStorage.getItem('jwt')
    if (jwtToken == null) {
        return <Navigate to={path} />
    }

    const decodedToken = jwt(jwtToken)

    // seconds elapsed since January 1, 1970 UTC
    const now = Math.round(new Date().getTime() / 1000)

    // check not expired (we consider it expired 5 minutes early) and check issuer
    if (decodedToken.exp - 300 < now || decodedToken.iss !== 'whosn-core') {
        localStorage.removeItem('jwt')
        return <Navigate to={path} />
    }

    return children
}

export default PrivateRoute
