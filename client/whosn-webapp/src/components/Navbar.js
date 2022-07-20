import '../styles/Navbar.css'
import { Link } from 'react-router-dom'

const unsetToken = () => {
    localStorage.removeItem('jwt')
}

const Navbar = () => {
    return (
        <section className="navbar">
            <Link to="/login" className="navbar-item" onClick={unsetToken}>
                Logout
            </Link>
            <Link to="/" className="navbar-item">
                Home
            </Link>
        </section>
    )
}

export default Navbar
