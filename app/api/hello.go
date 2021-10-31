package api

import (
	"sea/app/service"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	service.Water.ReGenWaterID()
	w, _ := service.Water.GetSelfWater()
	k, _ := crypto.GenerateKey("name", "email", "rsa", 4096)
	ks, _ := k.Armor()
	// fmt.Println(ks)
	// 	ks := `-----BEGIN PGP PRIVATE KEY BLOCK-----
	// Version: GopenPGP 2.2.4
	// Comment: https://gopenpgp.org

	// xcZYBGFre3ABEAC8q1aOwXpMIE9vPexn97SpcxXQ8A+EqwBBfGyh8FyVAdqC1d89
	// ZSQ+Mks2RhYif1XnJPHQCWw1r2t5qKS9FDypSNWP+cCjYV8sEriDajBTMgm4feJe
	// 0B1sk+3y1kpy3tSmNtxCvlxkDBFYktIDOPeaN+vvBBkYivLDRkKn9f72hCsQg145
	// vqx3FroMLm2bYVtfCTv9UdFL+wRCb1M47RlNucbWzZhJweOBBDk9ssEVPPUYoqLo
	// ByaE53t3r4UBTovJQI3mjG2PgYGF43Qg8QHnt0/z4AjwKABfbMtVKn4LmcUqjnir
	// /LwRQRQ11nQF6ru3/4o1n/MvyZeEqsgiJdyqSUReim3skse9GnRs4ng8IaHcCEOf
	// Pp+E9qmYUXShGDOJNs0ZJJI/iTp+Jcwy3kQEEYad2hF6iHkDnZwiWBAIoitqOdOe
	// mETKyvU53REYBrtO6Ja9Kq9d2J+5s+yj8wEb991pZScB20ih8PTErPNRgWN3OOWt
	// fHrR6nIMJh0imtHVB+t18j8wmBmNIIhWGOI91MroSGMwD0RY5D0kT/y1p1YPZ3wW
	// jnNjTf6zkcZmx4OKxzI3smtnapLBu4eLI2+h8la+BZGV3+BFVD9KLaWKL0HF5kTn
	// UZzqliS5uDA0ONxX8Kdy3cSWaHdVaD0hqYoyRUDsOG6TUXBDHgFMvSDRfwARAQAB
	// ABAAnJZ0E62fDDmw0/oIEM404oKC89OafZjTcIaNPuZYYM+WBQru6mzcZManR9dp
	// nZ4jtOEj1W9MoU4DnIiMmeWG51Vs8r7t1GwHanPrMh7bFuGr0bShQeHm2ev5DJgw
	// WCB7S6yyqK6Hzf16FBJfPY91aUjKh6jWD1mbXNkej3qK7x0JBP6TIMgRKhHxvhR2
	// ogRuqnTmljvmwvIGOgFdiBy46Jqs3jDrGSRT5wuTWpLW0AK0xYqOPJdar+nTvqUt
	// dcU0HYd3fobEk9t/nZkHFCQTWbb6I8WohXIdxRVtuTS4QnaQft6aZtwvVv61gG2K
	// BDMrM+5iTll/pATjvZdk3XqhJkJA9Ta56cRwK2q7pGdcNzarxwB4PzOUHfdgYKte
	// Sj1QAcF2Ht7LK8SitnRYwRQo/md4xT9ZB/qoUEzIQFZoghESgWOYsn94A0MF5ST9
	// gvvXlJSoO1IZJfwdcbaGvv58H6vsfMr5hG/i8P/2hZPE8ERlHoWXwqTmet4GmXq6
	// Y9PkQARAgehxjkO9aPrZDf8ggN3rv4X4Y01FlXla7EBX1F6p8nRV0WydwnKj5/yD
	// nLkQdSwDUWog1HRTunUw6YOIRabymGiOytPbXjvVOIupaOTIui8UphM4JmxH99Qx
	// BRaEbo+LZKRp5BT099T4+bsRsaaFxBHoFit3tO7Et7eQdakIAOw4AnI8+7ZiRV4V
	// O/xB2cTPBzOmvFhk+ereogjbuF0imyQ8noG0svnu9gishBdzqufGcGcHAGjuiHsI
	// sbe7hKg22w+qHcdff3RwvPSGEfdFWMNLYjV9m9ECQzO9+r9F1TA1DRpkLZMM0r0G
	// QklbP5+doFQmEakfmLZGmhSQxXPxc5kDJhSykaQr5j0TYQixS/jiQdr0gjIrGJdK
	// A3/MooT4wj05Q8szeI6lAhyWUrbwU1zqY4DVFBTvK441CcTOyYdkz6ByHK570R3C
	// PSD7RM1hLdC6pHzph2ji6QJXgKM/NKEgkBUbPfqZy1ahhnpS5w//QtRLdcs1lCd/
	// fLY1ybUIAMx3+eGIvPJbbLy6sG+H3W2p1a5lHbGPxALVsoQcbKHFtecwoT5xYcTu
	// wJT5Nu+LIJf3Zmu9Zu1J1AMN0nsA8xk6K7TrG9xtywn3PGutRd1Vg1M4nTpxhC44
	// a4yBo3pXqbXG5vnXSByURLNmVX9HpSyywnzRh6aGxg/XxDzSAzgOB31pD9kXDxVi
	// Mf4D8iFYD3ctecdBjoi/3xfTYoAceTwF0kXMIt6qAqUnAsKnXi8dV2OJ4LJdb3qE
	// qBceuwzeQmz3uSZYy+9FIIyje3zeNY7mwBSt+SC82WDRzONOoqcU9JDgb3KSl4jS
	// 7CgvjlhNN8m9UTFwKjXd1PwBgu703uMH/RkTWx0EC2kUBeX9bklNEiaLhTYuKEkE
	// X5q6y7Q3AuXPZo3/lNrMb0nMXtVsTlK3mNbcsb4r1QyTyq6n16vZ3uJfWQW1vNXE
	// qkMdcWiOmf0Np8evKdVqBwbAEcIELPZGj0FPk0/Y50lBdOpR1q5drUd5ua8Qnj0G
	// 8+1kmsQELI38dQ/0WFc0rpKCM0gtWz6EgwiSWK/B7aLCYlklnVKgw1jpbOyozaUz
	// q8mwX3SsoExw8X2HIzLhxAw917cSQXS8qTBD66QWjPsSLfxAk8ESfV4qcPY/3QVN
	// /0EvNEKXmdDiMXMD5dPZ5sYF/uNlzt6VjedFGpYycBaMlzmTbXp+xYCHt80MbmFt
	// ZSA8ZW1haWw+wsGKBBMBCAA+BQJha3twCZDjcYZQ/fbTLxahBGnpZXDBw/x8M9p4
	// ZuNxhlD99tMvAhsDAh4BAhkBAwsJBwIVCAMWAAICIgEAABDAD/kB5x7bTc78n3HF
	// SOVUc8oSlMBCeqZXQZJrHC8DlUVjbF7gO2EgvEJsBAHeVJYD9MZIyWQvzpUVI9hB
	// laJRbYNYPN/ApzfdgrfzmmINng2IJuwrL0Bs1FRlE9+mjdBBb9httHSyaHDZ7g+f
	// o1LDjtwtEZsceZU0a+3g9FrggjnJdHW+6ciB0zj4kmeER0VJfLRb3i7kcKhc9ob6
	// cPa/J5rSkonaL6wmVz4EbE7LYek0RIMNE8ckb3ys3XDqMcWkPX0I7CLirgsA4nGk
	// HaQY9eibnk/WYRWv/4OPdBfGAUGVx9CueGarCOeAjJ8fu0qF9SUMKodL1dqxDGNJ
	// G4gnQBmRYAs6o+wgA+dxAPd1iXl3RTn+/MouqZ1GRyFOg2gsIxl4b8Hm3lYdL4Wj
	// abMsyylqLZ/B8ebI1g3VvNSeQ/oPLLlcJXNdrijCYrK/Y6XsnZxWoaUcdcBfuv9n
	// aBBfozAK9Zefenzn74Bz9ytzhOVAeR/V+yx3uYv9OU8q+lVRI3AogkJeo8XyTJB9
	// Sj38VIQv9u/9PqG20ma/k+tHCCq172XyzpBiySKcTOTS0SVjRUxmcZTBDk7lutTu
	// x8RRD1YRcfYua8XvtNyizh2VruKSUOQU0uRjWIpUUrsRZCK3pWCJIH7p0LZb8QF4
	// I/0VnG64JZ6QzpQR+zOdjAcfrE+bOMfGWARha3twARAA4K2XthS6/yA0FTYN/mYe
	// Tcmd3gYRF1dl+47FgqH2T1Rc+mFI7kUeBiAydtZccIhuyDiDnKua66PM0z3kM4hm
	// U+T5TXmNcMjESD7kxSOwo3Mih0W3esx/EM6sWSZo3fby73WHTBba40RO8krM6QXc
	// vLMZtfI5RMsO5HnBQ7uFS94qg2t0HlVRYQP7KiR/oSVbk8J4QcT3OKwsFuV3nWPd
	// efsPSg10CQEAiJgMtVuBIWgNGaJY3MUQhe7/Z7gAZ4qsv5K91wxbKGKcccvkZUKL
	// FcdX0jO5NR95rDSBzfH+dVSXD0gxA5zLDGTKuivVFW0CS7PgaHbmbXBB+dAZ0YqC
	// W52+EDzetHyZgN44I5vzzWuryy5d6bRcW611r7kba8QFjSDeym0vvmgEOUFsU4Gn
	// MoyRcM1fvzrlcRzb8AvMmpF1bCUD7bdYey44KfaVa4q8wacM7klSQK0PwvSZ5Adt
	// 2NYXad18naI+oyzMrqgLTB04uOfKfhscKDqZHDlp0UwyGITRbKSujgCQiiEPoghq
	// mi/WltU97fH1S7eeOJlsoa5faGWXZLndBQY/T3l83Xo1tPPS9dCz69KdgsoZVoFh
	// AZOcOnsq3ffVMHi9GHW8HIds3Sr1Tvx7VJ1Md5MMYisUBnz4ewMFJrYsm18e2UJ0
	// uYFFcfzncgDq8C9RlhStiTcAEQEAAQAP/3imn5aGkVhZeb5+NYuMWO+auiy1A06O
	// f1RT+dpJkRROkFNgFKLm+NV/SrooKKHBkgPw3TQvrtSlNGOnk7jcoV6h0gJ5qKuN
	// FGJba8w1ezB6rmsH1BycXguvgDurdZrkdtaCPXJtv4NVdjOEWX80x7GJV+Atv/Ui
	// OdAVVeRCebUhC6Fg/Y0gcb/wPFsIGIV3ozK4to1S8RM82bw94zrgmtlbKX2xtcL0
	// B85SIFjvV2uSDZ7bYG7FYDcb/Yk79POlM1j4uvsVRopsUeiF/hMmZbv/PTjCz80o
	// JtWDqMToE6X8E9aABNoUw7DXG0pYKWgFDruKT5KpbtoL81HsH0E36QG5zVNvcrf6
	// abV156yatL4YYq8dvH3jdCcgOxvxZuC9aDf/PGgSgHXhr7XavvtY+Xm+TiRs+zv8
	// pgCDTYPwWmy/6i9zL7S261MRkddAelAL+2N2L4GkOzRWY61DscUg1plwOuKI4Vz9
	// euls5uxjK1o8O2W62FkK6MpRfO5gJDJU+sr4duF0PgcjQ4Aek27jwoHgYfz7RRdI
	// z3wz8CXvRQIZIosSiLfYSNaORA/gb+LXDARtUrjHMSIHvspUwIZFH9uRAyD0Bf8f
	// bMN/cLdHUTmIHsYX1Apo5P1F+88lrXOd0d3u8lawco5ZbSw/+rWKfhsZpRBExs9J
	// 5VurpysT4uiZCADj6vQdSD6n88+gKNL+o5Emq8zcL4NNLhE/+HF4iY34pfjeZp72
	// Wykhz5KkXVryS7RvCNscXpJXRghaCBYqyiKvuRF3UpRJY6wEvO3oKHOd98Ksmotc
	// c+PcgoyAOAnxNCCvNlKuiacE/OSNmEGnZd0v5dbFXKuvrDoC1ESURIFF474A19jB
	// p9wE6p3P0hZWdPQsvgv9VDHa6JCTDXl0iXcWp6XTUO29pjqXmCArOWvEY2HM+Lf9
	// B6m93TNrCyzTUbuYhKTIwsupl278EiMtfDBy4WX6GwGZptj0Efue99gopTSWa8mY
	// OzZuUYVeMaNLUzjzi7UOvMje/dYQyhgOt52TCAD8XHOplcBnjwjCxNS3z4rM8thp
	// iMFYXO+QLZUaJh7LW+6M9/P0pADdj/eFVNEepxLHfPhHA9wYg2cI69jJ7REhsNYn
	// poL0Brud+loh+1eQRnViFQuk9Q3IoUNKdCIYwkSjpL7n3djsGFN/GfA8N2Efqelt
	// wKOlFhXnWeJHh75HK0GMdwdXc/yoVRb3HABG5TgBbH2uu8AqDw8ErEItkrnTzhrU
	// EHzXRVNh4eTDfQ26KlEp11Lml/niwdKQeiP3CIK4m7/dOt5k4RkS5lbZuOSOZkdJ
	// OGJ9KNy8LzGX4gwC9dwHECtKSxX0KaKEid0WiTpCAA0Bb3O6SUxNv0IP1sxNCACf
	// jEicU5nf9yVs3I4rNQuUpaencCOrOfgYqUHvl1qUWeCOXauIw58dtnOlSJwT3GhZ
	// CSMGjVpO4O/X6Ayw93j5B1clLYEbHbOFLC3f6LM8zHZGu15JGRTle/vl1iLYaBzW
	// cmNN1hPkQ2NnCcI3IgphyTHTGIf6KdhjgRf+lg+Q5WVqtrN2J5xd1GMdnODkzfgY
	// lM8E0uoPNH0QVogRNSIY+ThuNczZJOkLD312cPY2hobR4Qz45CcCd4EXkmPsqaLC
	// oOldFsMAtL8+mCKebbSKIPDntZcT0F/DBqxjwasao68KQzPvUrfGyZV5ssSjKigw
	// 1uXxT0C9vzG5htc/r1buh0nCwXYEGAEIACoFAmFre3AJkONxhlD99tMvFqEEaell
	// cMHD/Hwz2nhm43GGUP320y8CGwwAAElbD/0QzNbjoPD+H4ptQVyzilxfUJDuKLQL
	// jFJuGg0DhMGMIT9SqLckkm7AYASoKA9xv2WUCbqUaUQnTUd8YMpZm0f2giH7xi7G
	// z2jvbw1RhUxGOv8xPwdmEh3c9T8B6nfSdGIn4DMAe/QzJPvCp3KYrY5gC5VM+o7t
	// 8SINSxeh2I96Ug7TF7CIMe5h1rqU/UFu/K8qe+5ajFoHn8TvXQcTXQedcA6dhqaK
	// 3VYhOB5BTUG3qvAidH5+tyDc3PkbBoCKIeoK1v5St6RqRqTxzI+JUPa4X3AluOIL
	// VxD3DqWfVR1I66mBFuJyoKwvHVGqp3HGMDKWQ9NzNN98PuCR8dOhJnuP8bc8md1C
	// aTgoIl3JtUOH1Q9OVfkC6C7/8lU0be6jB3lYv/AFtTqQmwVkMK+ShBL7Bh12GrbC
	// nlPsnvz/B5Fete+MoyC7zH7daJlXOkwnjk+923MzObg9yGHMJVflFwDVAevbCu6V
	// kg14zOZhTmCsNB7asOiEK3RP9xJB1CNlq9bZ11FEGqgGd0qaGPqcsETsIG5m1bog
	// tjz1KgJ7Uz5GMoyDp24icwdPwKd8TQZT7XtsZFHkms5wOux7G6kaBbAtn4/fkoj7
	// +rtJ4tyHSY7YqKVa2jjWtBsBEjxX9+WLpsB4aHC+VkWymjn8C33EH4UmqLpVBbfB
	// BlPhe0ILIBEOcQ==
	// =pAhK
	// -----END PGP PRIVATE KEY BLOCK-----`
	armored, _ := helper.SignCleartextMessageArmored(ks, nil, w.WaterId)
	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   gconv.String(armored),
	})
}
