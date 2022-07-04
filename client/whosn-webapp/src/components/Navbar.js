import '../styles/Navbar.css'

const Navbar = () => {
    return (
        <section className="navbar">
            <a href="/" className="navbar-item">
                Home
            </a>
            <a href="/profile" className="navbar-item">
                Profile
            </a>
        </section>
    )
}

export default Navbar
