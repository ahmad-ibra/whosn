import Navbar from './Navbar'
import '../styles/Header.css'
import logo from '../logo_small.png'
import { Link } from 'react-router-dom'

const Header = () => {
    return (
        <section className="header">
            <section className="header-top">
                <section className="header-top__logo">
                    <Link to="/" className="header-logo">
                        <img className="logo" src={logo} alt="WhosN? Logo" />
                    </Link>
                </section>
                <section className="header-top__navbar">
                    <section className="header-top__navigation">
                        <Navbar />
                    </section>
                    <hr className="header-top__seperator" />
                </section>
            </section>
        </section>
    )
}

export default Header
