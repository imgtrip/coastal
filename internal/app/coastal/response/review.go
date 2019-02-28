package response

import (
    "coastal/config/constant"
    "coastal/internal/app/coastal/types"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
)

func review(review types.ReviewWithImageRelation) *pb.Review {
    return &pb.Review{
        Id:               review.ID,
        CreatorId:        review.CreatorId,
        ReviewCategoryId: review.ReviewCategoryId,
        AgreeCount:       review.AgreeCount,
        DisagreeCount:    review.DisagreeCount,
        EndAt:            review.EndAt.Format(constant.DateTimeFormat),
        CreatedAt:        review.CreatedAt.Format(constant.DateTimeFormat),
        UpdatedAt:        review.UpdatedAt.Format(constant.DateTimeFormat),
        User:             user(review.Creator),
        ReviewCategory:   reviewCategory(review.ReviewCategory),
        Image:            image(review.ImageWithRelation.Image, review.ImageWithRelation.Tags),
        Closed:           review.Closed,
        Content:          review.Content,
        Opinion:          review.ReviewLog.Opinion,
    }
}

func reviewCategory(reviewCategory model.ReviewCategory) *pb.ReviewCategory {
    return &pb.ReviewCategory{
        Id:   reviewCategory.ID,
        Name: reviewCategory.Name,
    }
}

func reviewItems(reviews []types.ReviewWithImageRelation) []*pb.Review {
    var items []*pb.Review
    for _, r := range reviews {
        items = append(items, review(r))
    }
    return items
}

func Reviews(reviews []types.ReviewWithImageRelation, total int) *pb.ReviewsRes {
    return &pb.ReviewsRes{Items: reviewItems(reviews), Total: uint64(total)}
}
