import React from 'react';
import { isiContent } from '../../../Api/dataForumStatic';
import CardForum from './CardForum';
import { Link } from 'react-router-dom';
function ContentForum() {
  return (
    <>
      <Link to="">
        <div className="my-3 flex flex-col items-center gap-4">
          {isiContent.map((data, index) => (
            <CardForum
              key={index}
              id={data.id}
              title={data.title}
              content={data.content}
              date={data.date}
              answer={data.answer}
              like={data.like}
              profileImg={data.profileImg}
              user={data.user}
            />
          ))}
        </div>
      </Link>
    </>
  );
}

export default ContentForum;
