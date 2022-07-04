import '../styles/Navbar.css'
import { Link } from 'react-router-dom'

const Navbar = () => {
    return (
        <section className="navbar">
            <Link to="/" className="navbar-item">
                Home
            </Link>
            <Link to="/profile" className="navbar-item">
                Profile
            </Link>
        </section>
    )
}

export default Navbar
