package arknights

// CountPromotionMaterials returns upgrade/promotion materials required for the plan.
func CountPromotionMaterials(plans ...Plan) Materials {
	res := make(Materials)

	for _, plan := range plans {
		op := find(plan.Name)

		if plan.Promotion < 2 {
			res = res.Add(op.elite2)
		}

		if plan.Promotion < 1 {
			res = res.Add(op.elite1)
		}
	}

	res = res.Add(countChips(plans...).materials())

	return res
}
