import React from 'react';
import './App.css';
import './output.css';
import Spiner from './Assets/Spinners/Spiner';
import { isiContent } from '../src/Api/dataForumStatic';
import { isiBlog } from '../src/Api/dataBlogStatic';

//Routing
import { Routes, Route } from 'react-router-dom';
import AboutPage from './Component/pages/AboutPage';

const HomePage = React.lazy(() =>
  import('./Component/pages/HomePage')
);
const DetailForumPage = React.lazy(() =>
  import('./Component/pages/DetailForumPage')
);
const NotFound = React.lazy(() => import('./Component/pages/NotFound'));
const DetailBlogPage = React.lazy(() =>
  import('./Component/pages/DetailBlogPage')
);
const CreateBlogPage = React.lazy(() =>
  import('./Component/pages/CreateBlogPage')
);
const BlogPage = React.lazy(() => import('./Component/pages/BlogPage'));
const ForumPage = React.lazy(() => import('./Component/pages/ForumPage'));
const CreateForumPage = React.lazy(() =>
  import('./Component/pages/CreateForumPage')
);
const FaqPage = React.lazy(() => import('./Component/pages/FaqPage'));

function App() {
  return (
    <>
      <Routes>
        <Route
          path="/"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <HomePage />
            </React.Suspense>
          }
        />

        {/* forumPage */}
        <Route path="forumin">
          <Route
            index
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <ForumPage data={isiContent} />
              </React.Suspense>
            }
          />
          <Route
            path="createforum"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <CreateForumPage />
              </React.Suspense>
            }
          />
          <Route
            path="detailforum/:idforum"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <DetailForumPage data={isiContent} />
              </React.Suspense>
            }
          />
        </Route>

        {/* blogPage */}
        <Route path="blogin">
          <Route
            index
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <BlogPage data={isiBlog}/>
              </React.Suspense>
            }
          />
          <Route
            path="createblog"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <CreateBlogPage />
              </React.Suspense>
            }
          />
          <Route
            path="detailblog/:idblog"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <DetailBlogPage data={isiBlog} />
              </React.Suspense>
            }
          />
        </Route>

        <Route
          path="aboutpage"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <AboutPage />
            </React.Suspense>
          }
        />

        <Route
          path="faqpage"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <FaqPage />
            </React.Suspense>
          }
        />

        <Route
          path="*"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <NotFound />
            </React.Suspense>
          }
        />
      </Routes>
    </>
  );
}

export default App;
