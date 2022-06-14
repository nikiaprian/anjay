import React from 'react';
import ScrollButton from '../atoms/ScrollButton';
import ContentForum from '../molecules/CodeForumIn/ContentForum';
//eslint-disable-next-line
import Footer from '../molecules/Footer';
import Head from '../molecules/Head';
import Navbar from '../molecules/Navbar';

function ForumPage() {
  return (
    <>
      <div className="w-screen h-screen">
        <Navbar />
        <div className=" mt-40 w-10/12 md:w-8/12 mx-auto flex flex-col gap-4">
          <Head
            titleHead="Code ForumIn"
            contentHead="Tempat berkomunikasi para programmer dengan cara mengajukan dan menjawab sebuah pertanyaan"
            nameButton="Create Pertanyaan"
            path="/forumin/createforum"
          />
          <ContentForum />
        </div>
        <ScrollButton />
        <Footer />
      </div>
    </>
  );
}

export default ForumPage;
