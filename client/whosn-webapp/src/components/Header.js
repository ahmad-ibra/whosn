import Navbar from './Navbar'
import '../styles/Header.css'
import logo from '../logo_small.png'

const Header = () => {
    return (
        <section className="header">
            <section className="header-top">
                <section className="header-top__logo">
                    <a href="/" className="header-logo">
                        <img className="logo" src={logo} alt="WhosN? Logo" />
                    </a>
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
