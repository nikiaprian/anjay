import React from 'react';
import { isiBlog } from '../../../Api/dataBlogStatic';
import CardBlog from './CardBlog';
function ContentBlog() {
  return (
    <>
        <div className="my-3 flex flex-col items-center gap-4">
          {isiBlog.map((data, index) => (
            <CardBlog
              key={index}
              id={data.id}
              title={data.title}
              content={data.content}
              date={data.date}
              answer={data.answer}
              like={data.like}
              img={data.img}
              profileImg={data.profileImg}
              user={data.user}
            />
          ))}
        </div>

    </>
  );
}

export default ContentBlog;
