import React, { useContext, useState } from "react";
import { Link } from "react-router-dom";
import { Avatar, Button, Modal, Tag } from "antd";
import { FaRegHeart, FaRegComment } from "react-icons/fa";
import { FcLike } from "react-icons/fc";
import { AppContext } from "../../contexts/AppContext";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getUser } from "../../api/users";
import { getPostLikes, toggleLike } from "../../api/likes";
import { getPostComments } from "../../api/comments";
import { getPostCategories } from "../../api/categories";
import { displayTime } from "../../utils";
import { AlertBox } from "../../components/AlertBox";
import { Post } from "../../typings/types";

const BlogCard = ({ post }) => {
  const { accountOwner } = useContext(AppContext);
  const [open, setOpen] = useState(false);
  const accessToken = localStorage.getItem("accessToken");
  const queryClient = useQueryClient();
  const authorId = post.user_id;
  const postId = post.id;

  const { data: postOwner, isLoading } = useQuery({
    queryKey: ["users", authorId],
    queryFn: () => getUser(authorId),
  });

  const postLikesQuery = useQuery({
    queryKey: ["posts", postId, "likes"],
    queryFn: () => getPostLikes(postId),
  });

  const postCommentsQuery = useQuery({
    queryKey: ["posts", postId, "comments"],
    queryFn: () => getPostComments(postId),
  });

  const postCategoriesQuery = useQuery({
    queryKey: ["posts", postId, "categories"],
    queryFn: () => getPostCategories(postId),
  });

  const isLike = postLikesQuery.data?.some(
    (item) => item.user_id === accountOwner?.id && item.post_id === post.id
  );

  const toggleLikeMutaiton = useMutation({
    mutationKey: ["toggleLike", postId],
    mutationFn: toggleLike,

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts", postId, "likes"] });
    },
  });

  const handleLike = async () => {
    if (!accessToken) return setOpen(true);
    const userId = accountOwner?.id;
    if (typeof userId === 'number') {
      toggleLikeMutaiton.mutate({ postId, userId });
    }
  };
  

  return (
    <div className="flex flex-col w-full items-start gap-3 p-3 rounded-md shadow-xl border-[1px] dark:border-[#383838]">
      <div className="relative w-full overflow-hidden aspect-[10/6]">
        <img
          src={post.photo}
          alt=""
          className="object-cover w-full h-full rounded-md"
        />
        <Link to={`/blog/${post.id}`}>
          <div className="absolute bottom-0 left-0 right-0 top-0 h-full w-full overflow-hidden bg-[hsla(0,0%,98%,0.15)] bg-fixed opacity-0 transition duration-300 ease-in-out hover:opacity-100"></div>
        </Link>
      </div>
      <div className="flex gap-2">
        {postCategoriesQuery.data?.map((item) => (
          <Tag key={item.id} color="blue">
            <Link to={`/blog/category/${item.id}`} onClick={() => window.scrollTo(0, 0)}>
              {item.name}
            </Link>
          </Tag>
        ))}
      </div>
      <Link className="w-full" to={`/blog/${post.id}`}>
        <h1 className="font-bold h-[72px] w-full overflow-y-scroll custom-scrollbar">
          {post.title}
        </h1>
      </Link>
      {isLoading ? (
        <div>loading....</div>
      ) : (
        <div className="flex items-center gap-3">
          <Avatar src={postOwner?.profile} />
          <div className="flex justify-between gap-3 text-sm">
            <span>{postOwner?.name}</span>
            <span>, {displayTime(post)}</span>
          </div>
        </div>
      )}
      <div className="flex items-center w-full gap-5 p-2 ">
        <div
          onClick={handleLike}
          className="flex items-center gap-2 cursor-pointer"
        >
          {isLike ? <FcLike /> : <FaRegHeart />}
          <span>{postLikesQuery.data?.length || 0}</span>
        </div>
        <Link
          to={`/blog/${post.id}#comments`}
          className="flex items-center gap-2 cursor-pointer"
        >
          <FaRegComment />
          <span>{postCommentsQuery.data?.length || 0} </span>
        </Link>
      </div>
      <AlertBox
        open={open}
        setOpen={setOpen}
        text="Create an account to like for this blog."
      />
    </div>
  );
};

export default BlogCard;
