import React from 'react';
import Footer from '../molecules/Footer';
import Navbar from '../molecules/Navbar';

function AboutPage() {
  return (
    <>
      <Navbar />
      <div className="bg-white flex flex-col md:flex-row justify-center md:gap-28 items-center my-28 md:my-72">
        <p>Hai ini About</p>
      </div>
      <Footer />
    </>
  );
}

export default AboutPage;
