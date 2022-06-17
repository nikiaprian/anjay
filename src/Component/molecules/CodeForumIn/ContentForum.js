import React from 'react';
import CardForum from './CardForum';
function ContentForum(props) {
  const { dataApi, filter } = props;

  let temp;
  if (filter === '') {
    temp = dataApi;
  } else {
    let tempData = [];
    dataApi &&
      dataApi.map((item) => {
        return item?.tags.forEach((element) => {
          if (element.text.toLowerCase() === filter.toLowerCase()) {
            tempData.push(item);
          }
        });
      });
    temp = tempData;
  }

  return (
    <>
      <div className="my-3 flex flex-col items-center gap-4">
        {temp &&
          temp.map((data, index) => (
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
              tags={data.tags}
            />
          ))}
      </div>
    </>
  );
}

export default ContentForum;
